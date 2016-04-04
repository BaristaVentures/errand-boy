package bash

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/BaristaVentures/errand-boy/config"
	"github.com/BaristaVentures/errand-boy/routers/repos"
	"github.com/BaristaVentures/errand-boy/services/repotracker"
	"github.com/BaristaVentures/errand-boy/services/tracker"
	"github.com/BaristaVentures/errand-boy/utils"
	"github.com/Sirupsen/logrus"
)

func init() {
	repos.AddObserver("pr", pullRequestHandler)
}

func createCommandsFile(repoName string, repo *config.Repo, filepath string) error {
	if len(repo.Commands) == 0 {
		return errors.New("No deploy scripts for repo \"" + repoName + "\"")
	}
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	for _, cmd := range repo.Commands {
		_, err = file.WriteString(cmd + "\n")
		if err != nil {
			file.Close()
			os.Remove(filepath)
			return err
		}
	}
	if err = file.Close(); err != nil {
		os.Remove(filepath)
		return err
	}
	// Make it executable
	file, err = os.Open(filepath)
	if err != nil {
		os.Remove(filepath)
		return err
	}
	if err = file.Chmod(777); err != nil {
		os.Remove(filepath)
		return err
	}
	if err = file.Close(); err != nil {
		os.Remove(filepath)
		return err
	}
	return nil
}

var pullRequestHandler utils.ObserverFunc = func(payload interface{}) error {
	prPayload := payload.(repos.PullRequest)
	switch prPayload.Status {
	case "reopened":
		fallthrough
	case "opened":
		curConfig, err := config.Current()
		if err != nil {
			logrus.Error(err)
			return err
		}
		projectID, storyID, err := repotracker.GetTrackerData(&prPayload)
		if err != nil {
			logrus.Error(err)
			return err
		}
		project, err := curConfig.GetProject(projectID)
		if err != nil {
			logrus.Error(err)
			return err
		}

		for n, r := range project.Repos {
			name := n
			repo := r
			go func() {
				fileName := name + "-commands.sh"
				filePath := "./" + fileName
				err := createCommandsFile(name, repo, filePath)
				defer os.Remove(filePath)

				location := "/tmp"
				// Make sure to remove an existing script if present.
				err = exec.Command("ssh", repo.Host, "rm -f "+location+"/"+fileName).Run()
				if err != nil {
					logrus.Errorf("Error removing existent commands file on %s: %s", repo.Host, err.Error())
					// error might have been caused because rm didn't find a file to remove.
					// That's fine. Continue.
				}
				hostLocation := repo.Host + ":" + location
				logrus.Infof("Copying commands to remote host: scp %s %s", filePath, hostLocation)
				err = exec.Command("scp", filePath, hostLocation).Run()
				if err != nil {
					logrus.Errorf("Error copying %s to %s: %s", filePath, repo.Host, err.Error())
					return
				}

				command := "sh " + location + "/" + fileName
				logrus.Infof("Running commands on remote host: ssh %s %s", repo.Host, command)
				// Run the commands on the copied sh file like ssh host.com sh awesome-repo-commands.sh
				err = exec.Command("ssh", repo.Host, command).Start()
				if err != nil {
					logrus.Errorf("Error running commands for %s on %s: %s", name, repo.Host, err.Error())
					return
				}
				logrus.Infof("Successfully ran commands for %s on %s", name, repo.Host)
				trackerClient := tracker.New(os.Getenv(curConfig.TrackerAPIToken))
				comment := fmt.Sprintf("This story is ready to be tested on %s", repo.Host)
				_, err = trackerClient.CommentOnStory(projectID, storyID, comment)
				if err != nil {
					logrus.Error(err)
				}
			}()
		}
	}
	return nil
}
