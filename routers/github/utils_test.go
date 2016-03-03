package github

import (
	"testing"

	"github.com/BaristaVentures/errand-boy/services"
	"github.com/hooklift/assert"
)

func TestParseTrackerCodeOk(t *testing.T) {
	story := parseTrackerCode("[PT 13375p34k 1234]")
	assert.Equals(t, "13375p34k", story.ProjectID)
	assert.Equals(t, "1234", story.ID)
}

func TestParseTrackerCodeLong(t *testing.T) {
	story := parseTrackerCode("Solve all the project's issues. [PT 13375p34k 1234]")
	assert.Equals(t, "13375p34k", story.ProjectID)
	assert.Equals(t, "1234", story.ID)
}

func TestParseTrackerCodeInvalid(t *testing.T) {
	assert.Equals(t, (*trackerService.TrackerStory)(nil), parseTrackerCode("[PT 13375p34k"))
}

func TestParseTrackerCodeMissing(t *testing.T) {
	assert.Equals(t, (*trackerService.TrackerStory)(nil), parseTrackerCode("Nope."))
}
