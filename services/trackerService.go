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
	tracker = pivotal.NewClient(os.Getenv("PT_API_KEY"))
}

// SetStoryFinished sets the story with the given ID as finished.
func SetStoryFinished(projectID, ID int) (*pivotal.Story, error) {
	storyRequest := &pivotal.StoryRequest{State: pivotal.StoryStateFinished}
	s, _, err := tracker.Stories.Update(projectID, ID, storyRequest)
	if err != nil {
		return nil, err
	}
	return s, nil
}
