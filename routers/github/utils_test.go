package github

import (
	"testing"

	"github.com/hooklift/assert"
)

func TestParseTrackerCodeOk(t *testing.T) {
	projectID, storyID, err := parseTrackerCode("[PT 12321312 1234]")
	assert.Ok(t, err)
	assert.Equals(t, 12321312, projectID)
	assert.Equals(t, 1234, storyID)
}

func TestParseTrackerCodeLong(t *testing.T) {
	projectID, storyID, err := parseTrackerCode("Solve all the project's issues. [PT 123123 1234]")
	assert.Ok(t, err)
	assert.Equals(t, 123123, projectID)
	assert.Equals(t, 1234, storyID)
}

func TestParseTrackerCodeInvalid(t *testing.T) {
	_, _, err := parseTrackerCode("[PT 13375p34k")
	assert.Cond(t, err != nil, err.Error())
}

func TestParseTrackerCodeMissing(t *testing.T) {
	_, _, err := parseTrackerCode("Nope.")
	assert.Cond(t, err != nil, err.Error())
}
