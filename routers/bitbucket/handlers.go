package bitbucket

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/BaristaVentures/errand-boy/utils"
)

var events = []string{"pr"}
var eventsSubs = make(map[string]*utils.Observers)

// PullRequestPayload represents the body of BitBucket's PR webhook.
type PullRequestPayload struct {
	PR PullRequest `json:"pullrequest"`
}

type PullRequest struct {
	Title string `json:"title"`
	State string `json:"state"`
	URLs  URL    `json:"links"`
}

type URL struct {
	HTML Link `json:"html"`
}

type Link struct {
	Href string `json:"href"`
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

func papayaHandler(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusInternalServerError)
}
