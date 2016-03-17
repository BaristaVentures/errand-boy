package repotracker

import (
	"errors"

	"github.com/BaristaVentures/errand-boy/routers/repos"
	"github.com/BaristaVentures/errand-boy/services/tracker"
	"github.com/BaristaVentures/errand-boy/utils"
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
	case "opened":
		projectID, storyID, err := getTrackerData(&prPayload)
		if err != nil {
			return err
		}
		// Set the story as finished.
		_, err = trackerService.SetStoryState(projectID, storyID, "finished")
		if err != nil {
			return errors.New("Request to Pivotal Tracker API (update story) Failed.")
		}

		// Add a comment indicating the PR's URL.
		trackerService.CommentOnStory(projectID, storyID, "Check the PR @ "+prPayload.URL)
	}
	return nil
}
