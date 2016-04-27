package webhook

import (
	"fmt"
	"os"
	"testing"

	"github.com/BaristaVentures/errand-boy/config"
	"github.com/BaristaVentures/errand-boy/routers/repos"
	"github.com/BaristaVentures/errand-boy/testutil"
	"github.com/hooklift/assert"
)

func TestPRHandlerUnsupportedStatus(t *testing.T) {
	pr := repos.PullRequest{
		Status: "rejected by your mom",
	}
	err := pullRequestHandler(pr)
	assert.Cond(t, err == nil, "err should be nil if the PR status just doesn't trigger the hook.")
}

func TestPRHandlerInvalidTitle(t *testing.T) {
	pr := repos.PullRequest{
		Status: "opened",
		Title:  "Non-errand-boy-compliant title",
	}
	err := pullRequestHandler(pr)
	assert.Cond(t, err != nil, "err shouldn't be nil if the PR title doesn't have the expected format")
}

func TestPRHandlerEmptyHook(t *testing.T) {
	trackerID := 987654321
	repoName := "awesome-repo"
	conf := &config.Config{}
	reposMap := make(map[string]*config.Repo)
	// Configure a repo with no Hook
	reposMap[repoName] = &config.Repo{}
	projects := []*config.Project{
		&config.Project{
			TrackerID: trackerID,
			Repos:     reposMap,
		},
	}
	conf.Projects = projects
	configPath := "./test_eb-config.json"
	err := testutil.CreateConfigFile(conf, configPath)
	defer os.Remove(configPath)
	assert.Ok(t, err)

	_, err = config.Load(configPath)
	assert.Ok(t, err)

	pr := repos.PullRequest{
		Repo:   repoName,
		Status: "opened",
		Title:  fmt.Sprintf("[%d] %s", trackerID, "Just a PR Title"),
	}

	err = pullRequestHandler(pr)
	assert.Cond(t, err != nil, "err shouldn't be nil if the hook is empty.")
}

func TestGetHookForRepo(t *testing.T) {
	trackerID := 987654321
	repoName := "awesome-repo"
	hook := "http://webhook.com/build"
	conf := &config.Config{}
	reposMap := make(map[string]*config.Repo)
	reposMap[repoName] = &config.Repo{
		Hook: hook,
	}
	projects := []*config.Project{
		&config.Project{
			TrackerID: trackerID,
			Repos:     reposMap,
		},
	}
	conf.Projects = projects
	configPath := "./test_eb-config.json"
	err := testutil.CreateConfigFile(conf, configPath)
	defer os.Remove(configPath)
	assert.Ok(t, err)

	_, err = config.Load(configPath)
	assert.Ok(t, err)

	foundHook, err := getHookForRepo(trackerID, repoName)
	assert.Ok(t, err)
	assert.Equals(t, hook, foundHook)
}

func TestGetHookForRepoNoHook(t *testing.T) {
	trackerID := 987654321
	repoName := "awesome-repo"
	conf := &config.Config{}
	reposMap := make(map[string]*config.Repo)
	// Init a repo without a hook.
	reposMap[repoName] = &config.Repo{}
	projects := []*config.Project{
		&config.Project{
			TrackerID: trackerID,
			Repos:     reposMap,
		},
	}
	conf.Projects = projects
	configPath := "./test_eb-config.json"
	err := testutil.CreateConfigFile(conf, configPath)
	defer os.Remove(configPath)
	assert.Ok(t, err)

	_, err = config.Load(configPath)
	assert.Ok(t, err)

	foundHook, err := getHookForRepo(trackerID, repoName)
	assert.Ok(t, err)
	assert.Equals(t, "", foundHook)
}

func TestGetHookForRepoNoProject(t *testing.T) {
	trackerID := 987654321
	repoName := "awesome-repo"
	conf := &config.Config{}
	reposMap := make(map[string]*config.Repo)
	// Init a repo without a hook.
	reposMap[repoName] = &config.Repo{}
	projects := []*config.Project{
		&config.Project{
			TrackerID: trackerID,
			Repos:     reposMap,
		},
	}
	conf.Projects = projects
	configPath := "./test_eb-config.json"
	err := testutil.CreateConfigFile(conf, configPath)
	defer os.Remove(configPath)
	assert.Ok(t, err)

	_, err = config.Load(configPath)
	assert.Ok(t, err)
	// Call with a tracker ID that doesn't belong to any configured project.
	foundHook, err := getHookForRepo(00000000, repoName)
	assert.Cond(t, err != nil, "err shouldn't be nil when the specified tracker ID is not configured.")
	assert.Equals(t, "", foundHook)
}

func TestGetHookForRepoNoConfig(t *testing.T) {
	foundHook, err := getHookForRepo(00000000, "whatever-repo")
	assert.Cond(t, err != nil, "err shouldn't be nil when config.Load() hasn't been called.")
	assert.Equals(t, "", foundHook)
}
