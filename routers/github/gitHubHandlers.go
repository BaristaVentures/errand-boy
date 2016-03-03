package github

import (
	"github.com/BaristaVentures/errand-boy/services"
	"github.com/plimble/ace"
)

type pullRequestPayload struct {
	Action      string       `json:"action"`
	Number      string       `json:"number"`
	PullRequest *pullRequest `json:"pull_request"`
}

type pullRequest struct {
	Title    string `json:"title"`
	URL      string `json:"url"`
	MergedAt string `json:"merged_at"`
}

func pullRequestHandler(c *ace.C) {
	var prPayload pullRequestPayload
	c.ParseJSON(&prPayload)

	switch prPayload.Action {
	case "opened":
		projectID, storyID, ok := parseTrackerCode(prPayload.PullRequest.Title)
		if !ok {
			//FIXME: should this be a 400 tho
			c.AbortWithStatus(400)
		}
		// Set the story as finished.
		story, err := trackerService.SetStoryFinished(projectID, storyID)
		if err != nil {
			c.AbortWithStatus(500)
		}

		c.JSON(200, story)
		// Add a comment indicating the PR's URL.
		trackerService.CommentOnStory(projectID, storyID, "Check the PR @ "+prPayload.PullRequest.URL)
	}
}
