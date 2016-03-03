package trackerService

import (
	"os"

	"github.com/salsita/go-pivotaltracker/v5/pivotal"
)

var tracker *pivotal.Client

// TrackerStory maps a Tracker's story.
type TrackerStory struct {
	ID        int
	ProjectID int
}

func init() {
	tracker = pivotal.NewClient(os.Getenv("PT_API_TOKEN"))
}

// SetStoryFinished sets the story with the given ID as finished.
func SetStoryFinished(projectID, storyID int) (*pivotal.Story, error) {
	storyRequest := &pivotal.StoryRequest{State: pivotal.StoryStateFinished}
	s, _, err := tracker.Stories.Update(projectID, storyID, storyRequest)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// CommentOnStory adds a comment to story with the given ID.
func CommentOnStory(projectID, storyID int, comment string) (*pivotal.Comment, error) {
	pivotalComment := &pivotal.Comment{Text: comment, StoryId: storyID}
	updatedComment, _, err := tracker.Stories.AddComment(projectID, storyID, pivotalComment)
	return updatedComment, err
}
