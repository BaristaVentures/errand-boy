package github

import (
	"errors"

	"github.com/BaristaVentures/errand-boy/utils"
	"github.com/plimble/ace"
)

var events = []string{"pr"}
var eventsSubs = make(map[string]*utils.Observers)

type PullRequestPayload struct {
	Action string       `json:"action"`
	Number string       `json:"number"`
	PR     *PullRequest `json:"pull_request"`
}

type PullRequest struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

func init() {
	for _, ev := range events {
		observers := make(utils.Observers, 0)
		eventsSubs[ev] = &observers
	}
}

func AddObserver(event string, observer utils.Observer) error {
	_, ok := eventsSubs[event]
	if !ok {
		return errors.New("No such event: " + event)
	}
	eventsSubs[event].AddObserver(observer)
	return nil
}

func pullRequestHandler(c *ace.C) {
	var prPayload PullRequestPayload
	c.ParseJSON(&prPayload)
	// TODO: handle possible publisher errors.
	_ = eventsSubs["pr"].Publish(prPayload)
	c.String(200, "")
}
