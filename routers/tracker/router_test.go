package tracker

import (
	"fmt"
	"testing"

	"github.com/hooklift/assert"
)

func TestExtractDataFromComment(t *testing.T) {
	owner := "BaristaVentures"
	repo := "errand-boy"
	issueNo := 9
	comment := fmt.Sprintf("Check the PR @ https://github.com/%s/%s/pull/%d", owner, repo, issueNo)
	parsedOwner, parsedRepo, parsedIssueNo, err := extractDataFromComment(comment)
	assert.Ok(t, err)
	assert.Equals(t, owner, parsedOwner)
	assert.Equals(t, repo, parsedRepo)
	assert.Equals(t, issueNo, parsedIssueNo)
}

func TestExtractDataFromCommentNoURL(t *testing.T) {
	comment := "Check the PR @ "
	_, _, _, err := extractDataFromComment(comment)
	assert.Cond(t, err != nil, "Error shouldn't be nil when there's no URL.")
}

func TestExtractDataFromCommentInvalidURL(t *testing.T) {
	comment := "Check the PR @ %s"
	_, _, _, err := extractDataFromComment(comment)
	assert.Cond(t, err != nil, "Error shouldn't be nil when there's no valid URL.")
}

func TestExtractDataFromCommentURLNoHost(t *testing.T) {
	comment := "Check the PR @ 0890980"
	_, _, _, err := extractDataFromComment(comment)
	assert.Cond(t, err != nil, "Error shouldn't be nil when there's no host in the URL.")
}
