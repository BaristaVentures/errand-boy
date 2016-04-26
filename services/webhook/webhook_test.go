package webhook

import (
	"os"
	"testing"

	"github.com/BaristaVentures/errand-boy/config"
	"github.com/BaristaVentures/errand-boy/testutil"
	"github.com/hooklift/assert"
)

func TestWebhookPRHandler(t *testing.T) {
	t.Skip()
	pullRequestHandler(0)
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
