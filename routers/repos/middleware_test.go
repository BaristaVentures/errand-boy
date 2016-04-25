package repos

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/hooklift/assert"
)

func TestReplaceRequestBody(t *testing.T) {
	ghPayload := &gitHubPRPayload{
		Action: "opened",
		PR: &gitHubPR{
			Title:   "I <3 PRS [PT 1401024 114991501]",
			HtmlURL: "https://github.com/BaristaVentures/errand-boy/pull/9",
			Base: &gitHubRef{
				Repo: &gitHubRepo{
					Name: "lib-awesome",
				},
			},
		},
	}
	ghPayloadBytes, err := json.Marshal(ghPayload)
	assert.Ok(t, err)
	r, err := http.NewRequest("POST", "http://what.ever", bytes.NewReader(ghPayloadBytes))
	assert.Ok(t, err)

	genPayload := new(PullRequest)
	genPayload.HydrateFromGitHub(*r)

	genericPayloadBytes, err := json.Marshal(genPayload)
	assert.Ok(t, err)
	replaceRequestBody(genPayload, r)
	body, err := ioutil.ReadAll(r.Body)
	assert.Ok(t, err)
	spew.Dump(body)
	assert.Equals(t, genericPayloadBytes, body)
}
