package tracker

import (
	"fmt"
	"testing"

	"github.com/hooklift/assert"
)

func TestExtractDataFromComment(t *testing.T) {
	host := "github.com"
	owner := "BaristaVentures"
	repo := "errand-boy"
	issueNo := 9
	comment := fmt.Sprintf("Check the PR @ https://%s/%s/%s/pull/%d", host, owner, repo, issueNo)
	prData, err := extractDataFromComment(comment)
	assert.Ok(t, err)
	assert.Equals(t, host, prData.Host)
	assert.Equals(t, owner, prData.Owner)
	assert.Equals(t, repo, prData.RepoName)
	assert.Equals(t, issueNo, prData.Number)
}

func TestExtractDataFromCommentNoURL(t *testing.T) {
	comment := "Check the PR @ "
	_, err := extractDataFromComment(comment)
	assert.Cond(t, err != nil, "Error shouldn't be nil when there's no URL.")
}

func TestExtractDataFromCommentInvalidURL(t *testing.T) {
	comment := "Check the PR @ %s"
	_, err := extractDataFromComment(comment)
	assert.Cond(t, err != nil, "Error shouldn't be nil when there's no valid URL.")
}

func TestExtractDataFromCommentURLNoHost(t *testing.T) {
	comment := "Check the PR @ 0890980"
	_, err := extractDataFromComment(comment)
	assert.Cond(t, err != nil, "Error shouldn't be nil when there's no host in the URL.")
}
