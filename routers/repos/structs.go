package repos

import (
	"strings"

	"github.com/Sirupsen/logrus"
)

// PRConverter identifies service-specific Pull Request payloads (like GitHub's and BitBucket's),
// that can be converted into a generic PullRequest.
type PRConverter interface {
	ToGenericPR() *PullRequest
}

// PullRequest represents generic Pull Request info.
type PullRequest struct {
	Title  string
	Status string
	URL    string
}

// gitHubPRPayload represents the body of github's PR webhook.
type gitHubPRPayload struct {
	Action string    `json:"action"`
	PR     *gitHubPR `json:"pull_request"`
}

type gitHubPR struct {
	Title   string `json:"title"`
	HtmlURL string `json:"html_url"`
}

// bitBucketPRPayload represents the body of BitBucket's PR webhook.
type bitBucketPRPayload struct {
	PR *bitBucketPR `json:"pullrequest"`
}

type bitBucketPR struct {
	Title string        `json:"title"`
	State string        `json:"state"`
	URLs  *bitBucketURL `json:"links"`
}

type bitBucketURL struct {
	HTML *bitBucketLink `json:"html"`
}

type bitBucketLink struct {
	Href string `json:"href"`
}

// ToGenericPR transforms a GitHubPRPayload into a Generic one.
func (ghPayload *gitHubPRPayload) ToGenericPR() *PullRequest {

	genericPayload := &PullRequest{}
	genericPayload.Status = ghPayload.Action
	genericPayload.Title = ghPayload.PR.Title
	genericPayload.URL = ghPayload.PR.HtmlURL

	return genericPayload
}

// ToGenericPR transforms a BitBucketPRPayload into a Generic one.
func (bbPayload *bitBucketPRPayload) ToGenericPR() *PullRequest {

	genericPayload := &PullRequest{}
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

// GetContext implements logging.Logger
func (bbPayload *bitBucketPRPayload) GetContext() logrus.Fields {
	fields := logrus.Fields{
		"state": bbPayload.PR.State,
		"title": bbPayload.PR.Title,
		"urls": logrus.Fields{
			"html": logrus.Fields{
				"href": bbPayload.PR.URLs.HTML.Href,
			},
		},
	}
	return fields
}

// GetContext implements logging.Logger
func (ghPayload *gitHubPRPayload) GetContext() logrus.Fields {
	fields := logrus.Fields{
		"action": ghPayload.Action,
		"pull request": logrus.Fields{
			"title": ghPayload.PR.Title,
			"url":   ghPayload.PR.HtmlURL,
		},
	}
	return fields
}
