package webhook

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"

	"github.com/BaristaVentures/errand-boy/config"
	"github.com/BaristaVentures/errand-boy/routers/repos"
	"github.com/BaristaVentures/errand-boy/services/repotracker"
	"github.com/BaristaVentures/errand-boy/utils"
	"github.com/Sirupsen/logrus"
)

func init() {
	repos.AddObserver("pr", pullRequestHandler)
}

var pullRequestHandler utils.ObserverFunc = func(payload interface{}) error {
	prPayload := payload.(repos.PullRequest)
	switch prPayload.Status {
	case "reopened":
		fallthrough
	case "opened":
		projectID, _, err := repotracker.GetTrackerData(&prPayload)
		if err != nil {
			logrus.Error(err)
			return err
		}
		hookURL, err := getHookForRepo(projectID, prPayload.Title)
		if len(hookURL) == 0 {
			if err == nil {
				err = errors.New("No webhook url specified for repo " + prPayload.Title)
			}
			logrus.Error(err)
			return err
		}
		err = postRequest(hookURL, prPayload.OriginalBody, prPayload.Headers)
		if err != nil {
			logrus.Error(err)
			return err
		}
	}
	return nil
}

func postRequest(hookURL string, body []byte, headers http.Header) error {
	req, err := http.NewRequest(http.MethodPost, hookURL, bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header = headers
	res, err := new(http.Client).Do(req)
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Server returned with status code %d", res.StatusCode)
	}
	return err
}

func getHookForRepo(projectID int, repoTitle string) (string, error) {
	conf, err := config.Current()
	if err != nil {
		return "", err
	}
	project, err := conf.GetProject(projectID)
	if err != nil {
		return "", err
	}
	for name, repo := range project.Repos {
		if name == repoTitle {
			return repo.Hook, nil
		}
	}
	return "", nil
}
