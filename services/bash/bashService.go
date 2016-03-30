package bash

import (
	"os/exec"
	"strings"

	"github.com/BaristaVentures/errand-boy/config"
	"github.com/BaristaVentures/errand-boy/routers/repos"
	"github.com/BaristaVentures/errand-boy/services/repotracker"
	"github.com/BaristaVentures/errand-boy/utils"
	"github.com/Sirupsen/logrus"
)

func init() {
	repos.AddObserver("pr", pullRequestHandler)
}

// Run runs a command with the given args.
func Run(command string, args ...string) ([]byte, error) {
	cmd := exec.Command(command, args...)
	return cmd.Output()
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
		curConfig, err := config.Current()
		if err != nil {
			return err
		}
		project, err := curConfig.GetProject(projectID)
		if err != nil {
			return err
		}

		for n, r := range project.Repos {
			name := n
			repo := r
			go func() {
				for i, script := range repo.Scripts {
					logrus.Infof(
						"Running \"%s\" (%d/%d) for repo %s of project %d",
						script,
						i+1,
						len(repo.Scripts), name, projectID,
					)
					splitScript := strings.Split(script, " ")
					output, err := Run(splitScript[0], splitScript[1:]...)
					if err != nil {
						logrus.Error(err)
					} else {
						logrus.Info(string(output))
					}
				}
			}()
		}
	}
	return nil
}
