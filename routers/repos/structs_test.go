package repos

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/hooklift/assert"
)

func TestHydrateFromGitHub(t *testing.T) {
	payload := &gitHubPRPayload{
		Action: "opened",
		PR: &gitHubPR{
			Title:   "I <3 PRS [114991501]",
			HtmlURL: "https://github.com/BaristaVentures/errand-boy/pull/9",
			Base: &gitHubRef{
				Repo: &gitHubRepo{
					Name: "lib-awesome",
				},
			},
		},
	}
	jsonBytes, err := json.Marshal(payload)
	assert.Ok(t, err)
	body := bytes.NewReader(jsonBytes)

	r, err := http.NewRequest("POST", "http://some.url", body)
	assert.Ok(t, err)

	genPayload := new(PullRequest)
	err = genPayload.HydrateFromGitHub(*r)
	assert.Ok(t, err)
	// Test that all basic data was mapped OK
	assert.Equals(t, payload.Action, genPayload.Status)
	assert.Equals(t, payload.PR.Title, genPayload.Title)
	assert.Equals(t, payload.PR.HtmlURL, genPayload.URL)
	assert.Equals(t, r.Header, genPayload.Headers)
}

func TestBitBucketPRToGenericPR(t *testing.T) {
	payload := &bitBucketPRPayload{
		PR: &bitBucketPR{
			Title: "I <3 PRS [114991501]",
			State: "OPEN",
			URLs: &bitBucketURL{
				HTML: &bitBucketLink{
					Href: "http://barista-v.com",
				},
			},
		},
		Repo: &bitBucketRepo{
			Name: "lib-awesome",
		},
	}
	jsonBytes, err := json.Marshal(payload)
	assert.Ok(t, err)
	body := bytes.NewReader(jsonBytes)

	r, err := http.NewRequest("POST", "http://some.url", body)
	assert.Ok(t, err)

	genPayload := new(PullRequest)
	err = genPayload.HydrateFromBitBucket(*r)
	assert.Ok(t, err)
	assert.Equals(t, "opened", genPayload.Status)
	assert.Equals(t, payload.PR.Title, genPayload.Title)
	assert.Equals(t, payload.PR.URLs.HTML.Href, genPayload.URL)
	assert.Equals(t, r.Header, genPayload.Headers)
}
