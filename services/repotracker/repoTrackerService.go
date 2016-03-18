package repotracker

import (
	"github.com/BaristaVentures/errand-boy/routers/repos"
	"github.com/BaristaVentures/errand-boy/services/logging"
	"github.com/BaristaVentures/errand-boy/services/tracker"
	"github.com/BaristaVentures/errand-boy/utils"
	"github.com/Sirupsen/logrus"
)

var trackerService tracker.Service

func init() {
	repos.AddObserver("pr", pullRequestHandler)
}

// SetTrackerService sets the tracker.Service instance to be used.
func SetTrackerService(service tracker.Service) {
	trackerService = service
}

var pullRequestHandler utils.ObserverFunc = func(payload interface{}) error {
	prPayload := payload.(repos.PullRequest)
	switch prPayload.Status {
	case "reopened":
		fallthrough
	case "opened":
		projectID, storyID, err := getTrackerData(&prPayload)
		logrus.Info(projectID, storyID)
		if err != nil {
			logrus.Error(err)
			return err
		}
		// Set the story as finished.
		_, err = trackerService.SetStoryState(projectID, storyID, "finished")
		if err != nil {
			logrus.Error(err)
			return err
		}

		// Add a comment indicating the PR's URL.
		trackerService.CommentOnStory(projectID, storyID, "Check the PR @ "+prPayload.URL)
		logging.Info(&prPayload, "Sucessfully processed Pull Request:")
	}
	return nil
}
