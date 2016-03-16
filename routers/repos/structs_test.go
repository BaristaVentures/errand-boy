package repos

import (
	"testing"

	"github.com/hooklift/assert"
)

func TestGitHubPRToGenericPR(t *testing.T) {
	payload := &gitHubPRPayload{
		Action: "opened",
		PR: &gitHubPR{
			Title:   "I <3 PRS [PT 1401024 114991501]",
			HtmlURL: "https://google.com",
		},
	}

	genPayload := payload.ToGenericPR()
	assert.Equals(t, payload.Action, genPayload.Status)
	assert.Equals(t, payload.PR.Title, genPayload.Title)
	assert.Equals(t, payload.PR.HtmlURL, genPayload.URL)
}

func TestBitBucketPRToGenericPR(t *testing.T) {
	payload := &bitBucketPRPayload{
		PR: &bitBucketPR{
			Title: "I <3 PRS [PT 1401024 114991501]",
			State: "OPEN",
			URLs: &bitBucketURL{
				HTML: &bitBucketLink{
					Href: "http://barista-v.com",
				},
			},
		},
	}

	genPayload := payload.ToGenericPR()
	assert.Equals(t, "opened", genPayload.Status)
	assert.Equals(t, payload.PR.Title, genPayload.Title)
	assert.Equals(t, payload.PR.URLs.HTML.Href, genPayload.URL)
}
