package tracker

import "github.com/salsita/go-pivotaltracker/v5/pivotal"

// NewService returns a default TrackerService
func NewService(apiToken string) Service {
	return &trackerService{apiToken}
}

// Service defines the available consumable PivotalTracker API endpoints.
type Service interface {
	SetStoryState(projectID, storyID int, state string) (*pivotal.Story, error)
	CommentOnStory(projectID, storyID int, comment string) (*pivotal.Comment, error)
}

type trackerService struct {
	apiToken string
}

// SetStoryFinished sets the story with the given ID as finished.
func (ts *trackerService) SetStoryState(projectID, storyID int, state string) (*pivotal.Story, error) {
	client := pivotal.NewClient(ts.apiToken)
	storyRequest := &pivotal.StoryRequest{State: state}
	s, _, err := client.Stories.Update(projectID, storyID, storyRequest)
	return s, err
}

// CommentOnStory adds a comment to story with the given ID.
func (ts *trackerService) CommentOnStory(projectID, storyID int, comment string) (*pivotal.Comment, error) {
	client := pivotal.NewClient(ts.apiToken)
	pivotalComment := &pivotal.Comment{Text: comment, StoryId: storyID}
	updatedComment, _, err := client.Stories.AddComment(projectID, storyID, pivotalComment)
	return updatedComment, err
}
