package repotracker

import (
	"errors"
	"testing"

	"github.com/BaristaVentures/errand-boy/routers/repos"
	"github.com/hooklift/assert"
	"github.com/salsita/go-pivotaltracker/v5/pivotal"
)

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

func TestPRHandler(t *testing.T) {
	SetTrackerService(&goodMockService{})
	prMock := &repos.PullRequest{Title: "[PT 123123 12]", URL: "https://google.com", Status: "opened"}
	err := pullRequestHandler(*prMock)
	assert.Ok(t, err)
}

func TestPRHandlerInvalidCode(t *testing.T) {
	prMock := &repos.PullRequest{Title: "[PT 123123 12", URL: "https://google.com", Status: "opened"}
	err := pullRequestHandler(*prMock)
	assert.Cond(t, err != nil, "Error shouldn't be nil when code is invalid. Got:\n%s", err.Error())
}

func TestPRHandlerAPICallFailed(t *testing.T) {
	SetTrackerService(&badMockService{})
	prMock := &repos.PullRequest{Title: "[PT 123123 12", URL: "https://google.com", Status: "opened"}
	err := pullRequestHandler(*prMock)

	assert.Cond(t, err != nil, "Error shouldn't be nil when the update story API call fails. Got:\n%s", err.Error())
}
