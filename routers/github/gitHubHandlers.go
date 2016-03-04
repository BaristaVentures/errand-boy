package github

import (
	"github.com/BaristaVentures/errand-boy/services"
	"github.com/plimble/ace"
)

var trackerService tracker.Service

type pullRequestPayload struct {
	Action      string       `json:"action"`
	Number      string       `json:"number"`
	PullRequest *pullRequest `json:"pull_request"`
}

type pullRequest struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

// SetTrackerService sets the tracker.Service instance to be used.
func SetTrackerService(service tracker.Service) {
	trackerService = service
}

func pullRequestHandler(c *ace.C) {
	var prPayload pullRequestPayload
	c.ParseJSON(&prPayload)

	switch prPayload.Action {
	case "opened":
		projectID, storyID, err := parseTrackerCode(prPayload.PullRequest.Title)
		if err != nil {
			//FIXME: should this be a 400 tho
			c.AbortWithStatus(400)
			return
		}
		// Set the story as finished.
		story, err := trackerService.SetStoryState(projectID, storyID, "finished")
		if err != nil {
			c.AbortWithStatus(500)
			return
		}

		c.JSON(200, story)
		// Add a comment indicating the PR's URL.
		trackerService.CommentOnStory(projectID, storyID, "Check the PR @ "+prPayload.PullRequest.URL)
	}
}
