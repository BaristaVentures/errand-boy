package github

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hooklift/assert"
	"github.com/plimble/ace"
	"github.com/salsita/go-pivotaltracker/v5/pivotal"
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

type goodMockService struct{}

func (ms *goodMockService) SetStoryState(projectID, storyID int, state string) (*pivotal.Story, error) {
	return &pivotal.Story{}, nil
}

func (ms *goodMockService) CommentOnStory(projectID, storyID int, comment string) (*pivotal.Comment, error) {
	return &pivotal.Comment{}, nil
}

type badMockService struct{}

func (ms *badMockService) SetStoryState(projectID, storyID int, state string) (*pivotal.Story, error) {
	return nil, errors.New("")
}

func (ms *badMockService) CommentOnStory(projectID, storyID int, comment string) (*pivotal.Comment, error) {
	return nil, errors.New("")
}

// The actual tests.
func TestPRHandler(t *testing.T) {
	ms := &goodMockService{}
	SetTrackerService(ms)
	recorder := httptest.NewRecorder()
	rw := &mockRwWrapper{ResponseWriter: recorder}
	prMock := &pullRequest{Title: "[PT 123123 1234]", URL: "https://google.com"}
	prPayloadMock := &pullRequestPayload{Action: "opened", PullRequest: prMock}
	body, _ := json.Marshal(prPayloadMock)
	request, _ := http.NewRequest("POST", "", bytes.NewReader(body))
	c := &ace.C{Request: request, Writer: rw}
	pullRequestHandler(c)
	assert.Equals(t, 200, recorder.Code)
}

func TestPRHandlerInvalidCodeFormat(t *testing.T) {
	ms := &goodMockService{}
	SetTrackerService(ms)
	recorder := httptest.NewRecorder()
	rw := &mockRwWrapper{ResponseWriter: recorder}
	prMock := &pullRequest{Title: "[PT 1234]", URL: "https://google.com"}
	prPayloadMock := &pullRequestPayload{Action: "opened", PullRequest: prMock}
	body, _ := json.Marshal(prPayloadMock)
	request, _ := http.NewRequest("POST", "", bytes.NewReader(body))
	c := &ace.C{Request: request, Writer: rw}
	pullRequestHandler(c)
	assert.Equals(t, 400, recorder.Code)
}

func TestPRHandlerTrackerServiceError(t *testing.T) {
	ms := &badMockService{}
	SetTrackerService(ms)
	recorder := httptest.NewRecorder()
	rw := &mockRwWrapper{ResponseWriter: recorder}
	prMock := &pullRequest{Title: "[PT 1234 123213]", URL: "https://google.com"}
	prPayloadMock := &pullRequestPayload{Action: "opened", PullRequest: prMock}
	body, _ := json.Marshal(prPayloadMock)
	request, _ := http.NewRequest("POST", "", bytes.NewReader(body))
	c := &ace.C{Request: request, Writer: rw}
	pullRequestHandler(c)
	assert.Equals(t, 500, recorder.Code)
}
