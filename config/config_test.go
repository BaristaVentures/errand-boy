package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/hooklift/assert"
)

func TestLoadConfig(t *testing.T) {
	trackerAPIToken := "asb1234basdasd"
	trackerProjectID := 123581321
	repoSource := "github"
	repoName := "awesome-repo"
	repoToken := "asdsad23edadsd1234812"
	configContentFmt := `{
    "tracker_api_token": "%s",
    "projects": [
      {
        "tracker_id": %d,
        "repos": [
          {
            "source": "%s",
            "name": "%s",
            "token": "%s"
          }
        ]
      }
    ]
  }`

	formattedContent := fmt.Sprintf(
		configContentFmt,
		trackerAPIToken,
		trackerProjectID,
		repoSource,
		repoName,
		repoToken,
	)

	configPath := "./test_eb-config.json"
	file, err := os.Create(configPath)
	defer os.Remove(configPath)
	assert.Ok(t, err)

	file.WriteString(formattedContent)
	file.Close()

	config, err := Load(configPath)
	assert.Ok(t, err)

	assert.Equals(t, trackerAPIToken, config.TrackerAPIToken)
	assert.Equals(t, 1, len(config.Projects))
	project := config.Projects[0]
	assert.Equals(t, trackerProjectID, project.TrackerID)
	repo := project.Repos[0]
	assert.Equals(t, repoName, repo.Name)
	assert.Equals(t, repoSource, repo.Source)
	assert.Equals(t, repoToken, repo.Token)
}

func TestLoadMissingConfig(t *testing.T) {
	configPath := "./missing_eb-config.json"

	_, err := Load(configPath)
	assert.Cond(t, err != nil, "Error shouldn't be nil.")
}
