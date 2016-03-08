package github

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/BaristaVentures/errand-boy/utils"
)

var events = []string{"pr"}
var eventsSubs = make(map[string]*utils.Observers)

// PullRequestPayload represents the body of github's PR webhook.
type PullRequestPayload struct {
	Action string       `json:"action"`
	Number string       `json:"number"`
	PR     *PullRequest `json:"pull_request"`
}

// PullRequest represents the PR data contained in a PullRequestPayload that is relevant for Errand Boy.
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

// AddObserver adds an observer to the list.
func AddObserver(event string, observer utils.Observer) error {
	_, ok := eventsSubs[event]
	if !ok {
		return errors.New("No such event: " + event)
	}
	eventsSubs[event].AddObserver(observer)
	return nil
}

func pullRequestHandler(res http.ResponseWriter, req *http.Request) {
	var prPayload PullRequestPayload
	json.NewDecoder(req.Body).Decode(&prPayload)
	// // TODO: handle possible publisher errors.
	_ = eventsSubs["pr"].Publish(prPayload)
	res.WriteHeader(http.StatusOK)
}
