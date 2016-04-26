package webhook

import (
	"bytes"
	"errors"
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
		conf, err := config.Current()
		if err != nil {
			logrus.Error(err)
			return err
		}
		project, err := conf.GetProject(projectID)
		if err != nil {
			logrus.Error(err)
			return err
		}
		hookURL := ""
		for name, repo := range project.Repos {
			if name == prPayload.Title {
				hookURL = repo.Hook
			}
		}
		if len(hookURL) == 0 {
			err = errors.New("No webhook url specified for repo " + prPayload.Title)
			logrus.Error(err)
			return err
		}
		req, err := http.NewRequest(http.MethodPost, hookURL, bytes.NewReader(prPayload.OriginalBody))
		if err != nil {
			logrus.Error(err)
			return err
		}
		req.Header = prPayload.Headers
		res, err := new(http.Client).Do(req)
		if err != nil || res.StatusCode != http.StatusOK {
			logrus.Error(err)
			return err
		}
	}
	return nil
}
