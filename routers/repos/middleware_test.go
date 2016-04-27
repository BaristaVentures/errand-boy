package repos

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/hooklift/assert"
)

func TestReplaceRequestBody(t *testing.T) {
	ghPayload := &gitHubPRPayload{
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
	// Marshal the struct into json.
	ghPayloadBytes, err := json.Marshal(ghPayload)
	assert.Ok(t, err)
	// Create a fake request.
	r, err := http.NewRequest("POST", "http://what.ever", bytes.NewReader(ghPayloadBytes))
	assert.Ok(t, err)
	// Initialize the PullRequest
	genPayload := new(PullRequest)
	genPayload.HydrateFromGitHub(*r)

	genericPayloadBytes, err := json.Marshal(genPayload)
	assert.Ok(t, err)

	replaceRequestBody(genPayload, r)
	// Check that the body was actually replaced.
	body, err := ioutil.ReadAll(r.Body)
	assert.Ok(t, err)
	assert.Equals(t, genericPayloadBytes, body)
}
