package repotracker

import (
	"fmt"
	"os"
	"testing"

	"github.com/BaristaVentures/errand-boy/config"
	"github.com/BaristaVentures/errand-boy/routers/repos"
	"github.com/BaristaVentures/errand-boy/testutil"
	"github.com/Sirupsen/logrus"
	"github.com/hooklift/assert"
)

func TestGetTrackerData(t *testing.T) {
	// We don't want logs while running tests.
	logrus.SetLevel(logrus.ErrorLevel)

	trackerAPIToken := "asb1234basdasd"
	trackerProjectID := 123581321
	repoName := "awesome-repo-1"
	repoToken := "asdsad23edadsd1234812"
	host := "localhost"
	port := 8000
	conf := &config.Config{
		TrackerAPIToken: trackerAPIToken,
		Projects: []*config.Project{
			&config.Project{
				TrackerID: trackerProjectID,
				Repos: map[string]*config.Repo{
					repoName: &config.Repo{
						Token: repoToken,
						Host:  host,
						Port:  port,
					},
				},
			},
		},
	}

	configPath := "./test_eb-config.json"
	err := testutil.CreateConfigFile(conf, configPath)
	defer os.Remove(configPath)
	assert.Ok(t, err)

	_, err = config.Load(configPath)
	assert.Ok(t, err)
	storyID := 123456

	pr := &repos.PullRequest{
		Repo:  repoName,
		Title: fmt.Sprintf("Awesome PR to solve everything [%d]", storyID),
	}
	parsedTrackerID, parsedStoryID, err := GetTrackerData(pr)
	assert.Ok(t, err)
	assert.Equals(t, trackerProjectID, parsedTrackerID)
	assert.Equals(t, storyID, parsedStoryID)
}

func TestGetTrackerDataNoCodeFormat(t *testing.T) {
	pr := &repos.PullRequest{
		Repo:  "a-repo",
		Title: "Bad PR Title with no code format :(",
	}
	_, _, err := GetTrackerData(pr)
	assert.Cond(t, err != nil, "Err shouldn't be nil when no code format is present in the PR title.")
}

func TestGetTrackerDataNoProjectConfig(t *testing.T) {
	pr := &repos.PullRequest{
		Repo:  "some-dudes-repo",
		Title: "Awesome PR with a code in the title but with no matching config :'(' [1234]",
	}
	curConfig, _ := config.Current()
	curConfig = &config.Config{}
	curConfig.Projects = []*config.Project{}
	_, _, err := GetTrackerData(pr)
	assert.Cond(t, err != nil, "Err shouldn't be nil when there's no matching config for that repo.")
}
