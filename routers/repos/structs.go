package repos

import (
	log "github.com/Sirupsen/logrus"
	"strings"
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

// debug method
func (genericPR *PullRequest) debug(logger *log.Entry) {
	logger.WithFields(log.Fields{
		"status": genericPR.Status,
		"title":  genericPR.Title,
		"url":    genericPR.URL,
	}).Info("Normalized fields")
}

func createLogContext(from string) *log.Entry {
	context := log.WithFields(log.Fields{
		"method": "ToGenericPR",
		"from":   from,
	})
	context.Info("Starting Normalization")
	return context
}

// ToGenericPR transforms a GitHubPRPayload into a Generic one.
func (ghPayload *gitHubPRPayload) ToGenericPR() *PullRequest {
	contextLogger := createLogContext("github")

	genericPayload := &PullRequest{}
	genericPayload.Status = ghPayload.Action
	genericPayload.Title = ghPayload.PR.Title
	genericPayload.URL = ghPayload.PR.HtmlURL

	genericPayload.debug(contextLogger)
	return genericPayload
}

// ToGenericPR transforms a BitBucketPRPayload into a Generic one.
func (bbPayload *bitBucketPRPayload) ToGenericPR() *PullRequest {
	contextLogger := createLogContext("bitbucket")

	genericPayload := &PullRequest{}
	// bbPayload.PR.State can be OPEN|MERGED|DECLINED
	if bbPayload.PR.State == "OPEN" {
		genericPayload.Status = "opened"
	} else {
		genericPayload.Status = strings.ToLower(bbPayload.PR.State)
	}
	genericPayload.Title = bbPayload.PR.Title
	genericPayload.URL = bbPayload.PR.URLs.HTML.Href

	genericPayload.debug(contextLogger)
	return genericPayload
}
