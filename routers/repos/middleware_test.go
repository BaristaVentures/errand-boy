package repos

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/Sirupsen/logrus"
	"github.com/hooklift/assert"
)

func TestReplaceRequestBody(t *testing.T) {
	logrus.SetLevel(logrus.ErrorLevel)
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
	r, err := http.NewRequest("POST", "what/ever", ioutil.NopCloser(bytes.NewBuffer(ghPayloadBytes)))
	assert.Ok(t, err)
	genericPayloadBytes, err := json.Marshal(ghPayload.ToGenericPR())
	assert.Ok(t, err)
	replaceRequestBody(ghPayload, r)
	body, err := ioutil.ReadAll(r.Body)
	assert.Ok(t, err)
	assert.Equals(t, genericPayloadBytes, body)
}
