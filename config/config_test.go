package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/Sirupsen/logrus"
	"github.com/hooklift/assert"
)

func TestLoadConfig(t *testing.T) {
	// We don't want logs while running tests.
	logrus.SetLevel(logrus.ErrorLevel)

	trackerAPIToken := "asb1234basdasd"
	trackerProjectID := 123581321
	repoName := "awesome-repo-1"
	repoSource := "github"
	repoToken := "asdsad23edadsd1234812"
	configContentFmt := `{
    "tracker_api_token": "%s",
    "projects": [
      {
        "tracker_id": %d,
        "repos": {
          "%s": {
            "source": "%s",
            "token": "%s"
          }
        }
      }
    ]
  }`

	formattedContent := fmt.Sprintf(
		configContentFmt,
		trackerAPIToken,
		trackerProjectID,
		repoName,
		repoSource,
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
	repo := project.Repos[repoName]
	assert.Equals(t, repoSource, repo.Source)
	assert.Equals(t, repoToken, repo.Token)
}

func TestLoadMissingConfig(t *testing.T) {
	configPath := "./missing_eb-config.json"

	_, err := Load(configPath)
	assert.Cond(t, err != nil, "Error shouldn't be nil.")
}
