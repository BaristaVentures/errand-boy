package github

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/BaristaVentures/errand-boy/utils"
	"github.com/hooklift/assert"
)

func TestPRHandler(t *testing.T) {
	called := false
	var sub utils.ObserverFunc
	sub = func(payload interface{}) error {
		called = true
		return nil
	}
	AddObserver("pr", sub)
	recorder := httptest.NewRecorder()
	prMock := &PullRequest{Title: "[PT 123123 1234]", URL: "https://google.com"}
	prPayloadMock := &PullRequestPayload{Action: "opened", PR: prMock}
	body, _ := json.Marshal(prPayloadMock)
	request, _ := http.NewRequest("POST", "", bytes.NewReader(body))
	pullRequestHandler(recorder, request)
	assert.Equals(t, true, called)
}
