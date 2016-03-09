package repos

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/BaristaVentures/errand-boy/utils"
)

var events = []string{"pr"}
var eventsSubs = make(map[string]*utils.Observers)

// GenericPRPayload represents generic Pull Request info.
type GenericPRPayload struct {
	Title  string
	Status string
	URL    string
}

// GitHubPRPayload represents the body of github's PR webhook.
type GitHubPRPayload struct {
	Action string    `json:"action"`
	PR     *GitHubPR `json:"pull_request"`
}

// GitHubPR represents the PR data contained in a PullRequestPayload that is relevant for Errand Boy.
type GitHubPR struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

// PullRequestPayload represents the body of BitBucket's PR webhook.
type BitBucketPRPayload struct {
	PR BitBucketPR `json:"pullrequest"`
}

type BitBucketPR struct {
	Title string       `json:"title"`
	State string       `json:"state"`
	URLs  BitBucketURL `json:"links"`
}

type BitBucketURL struct {
	HTML BitBucketLink `json:"html"`
}

type BitBucketLink struct {
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
	var prPayload GenericPRPayload
	json.NewDecoder(req.Body).Decode(&prPayload)
	// // TODO: handle possible publisher errors.
	_ = eventsSubs["pr"].Publish(prPayload)
	res.WriteHeader(http.StatusOK)
}

// ToGeneric transforms a GitHubPRPayload into a Generic one.
func (ghPayload *GitHubPRPayload) ToGeneric() *GenericPRPayload {
	genericPayload := &GenericPRPayload{}
	genericPayload.Status = ghPayload.Action
	genericPayload.Title = ghPayload.PR.Title
	genericPayload.URL = ghPayload.PR.URL
	return genericPayload
}

// ToGeneric transforms a BitBucketPRPayload into a Generic one.
func (bbPayload *BitBucketPRPayload) ToGeneric() *GenericPRPayload {
	genericPayload := &GenericPRPayload{}
	// bbPayload.PR.State can be OPEN|MERGED|DECLINED
	if bbPayload.PR.State == "OPEN" {
		genericPayload.Status = "opened"
	} else {
		genericPayload.Status = strings.ToLower(bbPayload.PR.State)
	}
	genericPayload.Title = bbPayload.PR.Title
	genericPayload.URL = bbPayload.PR.URLs.HTML.Href

	return genericPayload
}
