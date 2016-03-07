package github

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/BaristaVentures/errand-boy/utils"
	"github.com/hooklift/assert"
	"github.com/plimble/ace"
)

// mockRwWrapper is a mock struct that implements ace.ResponseWriter. We'll use it to be able to mock
// handler calls.
type mockRwWrapper struct {
	http.ResponseWriter
	status int
	size   int
}

// Mock methods
func (rw *mockRwWrapper) Status() int {
	return rw.status
}

func (rw *mockRwWrapper) Size() int {
	return rw.size
}

func (rw *mockRwWrapper) Written() bool {
	return rw.status != 0
}

func (rw *mockRwWrapper) Before(before func(ace.ResponseWriter)) {}

func (rw *mockRwWrapper) Flush() {}

// End mock response writer methods.

// The actual tests.
func TestPRHandler(t *testing.T) {
	called := false
	var sub utils.ObserverFunc
	sub = func(payload interface{}) error {
		called = true
		return nil
	}
	AddObserver("pr", sub)
	recorder := httptest.NewRecorder()
	rw := &mockRwWrapper{ResponseWriter: recorder}
	prMock := &PullRequest{Title: "[PT 123123 1234]", URL: "https://google.com"}
	prPayloadMock := &PullRequestPayload{Action: "opened", PR: prMock}
	body, _ := json.Marshal(prPayloadMock)
	request, _ := http.NewRequest("POST", "", bytes.NewReader(body))
	c := &ace.C{Request: request, Writer: rw}
	pullRequestHandler(c)
	assert.Equals(t, true, called)
}
