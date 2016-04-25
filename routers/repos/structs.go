package repos

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Sirupsen/logrus"
)

// PullRequest represents generic Pull Request info.
type PullRequest struct {
	OriginalBody []byte
	Headers      http.Header
	Title        string
	Status       string
	Repo         string
	URL          string
}

// gitHubPRPayload represents the body of github's PR webhook.
type gitHubPRPayload struct {
	Action string    `json:"action"`
	PR     *gitHubPR `json:"pull_request"`
}

type gitHubPR struct {
	Title   string     `json:"title"`
	HtmlURL string     `json:"html_url"`
	Base    *gitHubRef `json:"base"`
}

type gitHubRef struct {
	Repo *gitHubRepo `json:"repo"`
}

type gitHubRepo struct {
	Name string `json:"name"`
}

// bitBucketPRPayload represents the body of BitBucket's PR webhook.
type bitBucketPRPayload struct {
	PR   *bitBucketPR   `json:"pullrequest"`
	Repo *bitBucketRepo `json:"repository"`
}

type bitBucketRepo struct {
	Name string `json:"name"`
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

// HydrateFromBitBucket hydrates a *PullRequest from a BitBucket Request.
func (pr *PullRequest) HydrateFromBitBucket(r http.Request) error {
	bbPayload := new(bitBucketPRPayload)
	json.NewDecoder(r.Body).Decode(&bbPayload)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	pr.OriginalBody = body
	pr.Headers = r.Header
	// bbPayload.PR.State can be OPEN|MERGED|DECLINED
	if bbPayload.PR.State == "OPEN" {
		pr.Status = "opened"
	} else {
		pr.Status = strings.ToLower(bbPayload.PR.State)
	}
	pr.Title = bbPayload.PR.Title
	pr.URL = bbPayload.PR.URLs.HTML.Href
	pr.Repo = bbPayload.Repo.Name
	return nil
}

// HydrateFromGitHub hydrates a *PullRequest from a GitHub Request.
func (pr *PullRequest) HydrateFromGitHub(r http.Request) error {
	ghPayload := new(gitHubPRPayload)
	json.NewDecoder(r.Body).Decode(&ghPayload)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	pr.OriginalBody = body
	pr.Headers = r.Header
	pr.Status = ghPayload.Action
	pr.Title = ghPayload.PR.Title
	pr.URL = ghPayload.PR.HtmlURL
	pr.Repo = ghPayload.PR.Base.Repo.Name
	return nil
}

// GetContext implements logging.Logger
func (pr *PullRequest) GetContext() logrus.Fields {
	fields := logrus.Fields{
		"title":  pr.Title,
		"status": pr.Status,
		"url":    pr.URL,
		"repo":   pr.Repo,
	}
	return fields
}
