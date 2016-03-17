package repotracker

import (
	"fmt"
	"testing"

	"github.com/BaristaVentures/errand-boy/config"
	"github.com/BaristaVentures/errand-boy/routers/repos"
	"github.com/hooklift/assert"
)

func TestGetTrackerData(t *testing.T) {
	trackerID := 987654321
	storyID := 123456
	repoName := "awesome-repo"
	conf := config.Current()
	reposMap := make(map[string]*config.Repo)
	reposMap[repoName] = &config.Repo{}
	projects := []*config.Project{
		&config.Project{
			TrackerID: trackerID,
			Repos:     reposMap,
		},
	}
	conf.Projects = projects

	pr := &repos.PullRequest{
		Repo:  repoName,
		Title: fmt.Sprintf("Awesome PR to solve everything [%d]", storyID),
	}
	parsedTrackerID, parsedStoryID, err := getTrackerData(pr)
	assert.Ok(t, err)
	assert.Equals(t, trackerID, parsedTrackerID)
	assert.Equals(t, storyID, parsedStoryID)
}

func TestGetTrackerDataNoCodeFormat(t *testing.T) {
	pr := &repos.PullRequest{
		Repo:  "a-repo",
		Title: "Bad PR Title with no code format :(",
	}
	_, _, err := getTrackerData(pr)
	assert.Cond(t, err != nil, "Err shouldn't be nil when no code format is present in the PR title.")
}

func TestGetTrackerDataNoProjectConfig(t *testing.T) {
	pr := &repos.PullRequest{
		Repo:  "some-dudes-repo",
		Title: "Awesome PR with a code in the title but with no matching config :'(' [1234]",
	}
	conf := config.Current()
	conf.Projects = []*config.Project{}
	_, _, err := getTrackerData(pr)
	assert.Cond(t, err != nil, "Err shouldn't be nil when there's no matching config for that repo.")
}
