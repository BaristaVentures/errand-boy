package config

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/Sirupsen/logrus"
	"github.com/hooklift/assert"
)

func createConfigFile(conf *Config, configPath string) error {
	configBytes, err := json.Marshal(conf)
	if err != nil {
		return err
	}
	file, err := os.Create(configPath)
	if err != nil {
		return err
	}
	file.Write(configBytes)
	file.Close()
	return nil
}

func TestLoadConfig(t *testing.T) {
	// We don't want logs while running tests.
	logrus.SetLevel(logrus.ErrorLevel)

	trackerAPIToken := "asb1234basdasd"
	trackerProjectID := 123581321
	repoName := "awesome-repo-1"
	repoToken := "asdsad23edadsd1234812"
	conf := &Config{
		TrackerAPIToken: trackerAPIToken,
		Projects: []*Project{
			&Project{
				TrackerID: trackerProjectID,
				Repos: map[string]*Repo{
					repoName: &Repo{
						Token: repoToken,
					},
				},
			},
		},
	}

	configPath := "./test_eb-config.json"
	err := createConfigFile(conf, configPath)
	assert.Ok(t, err)
	defer os.Remove(configPath)

	loadedConf, err := Load(configPath)
	assert.Ok(t, err)

	assert.Equals(t, trackerAPIToken, loadedConf.TrackerAPIToken)
	assert.Equals(t, 1, len(config.Projects))
	project := config.Projects[0]
	assert.Equals(t, trackerProjectID, project.TrackerID)
	repo := project.Repos[repoName]
	assert.Equals(t, repoToken, repo.Token)
}

func TestLoadMissingConfig(t *testing.T) {
	configPath := "./missing_eb-config.json"

	_, err := Load(configPath)
	assert.Cond(t, err != nil, "Error shouldn't be nil.")
}

func TestCurrentMissingConfig(t *testing.T) {
	config = nil
	_, err := Current()
	assert.Cond(t, err != nil, "Error shouldn't be nil.")
}

func TestGetProject(t *testing.T) {
	trackerAPIToken := "asb1234basdasd"
	trackerProjectID := 123581321
	repoName := "awesome-repo-1"
	repoToken := "asdsad23edadsd1234812"
	conf := &Config{
		TrackerAPIToken: trackerAPIToken,
		Projects: []*Project{
			&Project{
				TrackerID: trackerProjectID,
				Repos: map[string]*Repo{
					repoName: &Repo{
						Token: repoToken,
					},
				},
			},
		},
	}

	configPath := "./test_eb-config.json"
	err := createConfigFile(conf, configPath)
	assert.Ok(t, err)
	defer os.Remove(configPath)

	loadedConf, err := Load(configPath)
	assert.Ok(t, err)
	project, err := loadedConf.GetProject(trackerProjectID)
	assert.Ok(t, err)

	assert.Equals(t, trackerProjectID, project.TrackerID)
	assert.Equals(t, len(conf.Projects[0].Repos), len(project.Repos))
}
