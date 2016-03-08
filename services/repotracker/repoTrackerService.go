package repotracker

import (
	"errors"
	"os"

	"github.com/BaristaVentures/errand-boy/routers/bitbucket"
	"github.com/BaristaVentures/errand-boy/routers/github"
	"github.com/BaristaVentures/errand-boy/services/tracker"
	"github.com/BaristaVentures/errand-boy/utils"
)

var trackerService tracker.Service

func init() {
	service := tracker.NewService(os.Getenv("PT_API_TOKEN"))
	SetTrackerService(service)
	github.AddObserver("pr", ghPullRequestHandler)
	bitbucket.AddObserver("pr", bbPullRequestHandler)
}

// SetTrackerService sets the tracker.Service instance to be used.
func SetTrackerService(service tracker.Service) {
	trackerService = service
}

var ghPullRequestHandler utils.ObserverFunc = func(payload interface{}) error {
	prPayload := payload.(github.PullRequestPayload)
	switch prPayload.Action {
	case "opened":
		projectID, storyID, err := parseTrackerCode(prPayload.PR.Title)
		if err != nil {
			return errors.New("Invalid Pivotal Tracker Code")
		}
		// Set the story as finished.
		_, err = trackerService.SetStoryState(projectID, storyID, "finished")
		if err != nil {
			return errors.New("Request to Pivotal Tracker API (update story) Failed.")
		}

		// Add a comment indicating the PR's URL.
		trackerService.CommentOnStory(projectID, storyID, "Check the PR @ "+prPayload.PR.URL)
	}
	return nil
}

var bbPullRequestHandler utils.ObserverFunc = func(payload interface{}) error {
	prPayload := payload.(bitbucket.PullRequestPayload)
	switch prPayload.PR.State {
	case "OPEN":
		projectID, storyID, err := parseTrackerCode(prPayload.PR.Title)
		if err != nil {
			return errors.New("Invalid Pivotal Tracker Code")
		}
		// Set the story as finished.
		_, err = trackerService.SetStoryState(projectID, storyID, "finished")
		if err != nil {
			return errors.New("Request to Pivotal Tracker API (update story) Failed.")
		}

		// Add a comment indicating the PR's URL.
		trackerService.CommentOnStory(projectID, storyID, "Check the PR @ "+prPayload.PR.URLs.HTML.Href)
	}
	return nil
}
