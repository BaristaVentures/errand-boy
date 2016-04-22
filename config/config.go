package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

var config *Config

// Config encapsulates Errand Boy's general config.
type Config struct {
	TrackerAPIToken string     `json:"tracker_api_token"`
	Projects        []*Project `json:"projects"`
}

// Project represents a project-specific config.
type Project struct {
	TrackerID int              `json:"tracker_id"`
	Repos     map[string]*Repo `json:"repos"`
}

// Repo represents a repository
type Repo struct {
	Token    string   `json:"token"`
	Host     string   `json:"host"`
	Port     int      `json:"port"`
	Commands []string `json:"commands"`
	Hook     string   `json:"hook"`
}

// GetProject returns a project if it matches
func (conf *Config) GetProject(trackerID int) (*Project, error) {
	if config == nil {
		return nil, errors.New("No loaded configuration.")
	}
	for _, p := range config.Projects {
		if p.TrackerID == trackerID {
			return p, nil
		}
	}
	return nil, fmt.Errorf("No project found with TrackerID: %d", trackerID)
}

// Load  parses the config from a json file to a *Config and returns it.
func Load(path string) (*Config, error) {
	config = &Config{}
	reader, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(bytes, config)
	return config, nil
}

// Current returns the current config.
func Current() (*Config, error) {
	// TODO: Maybe we should check the last time the config file was "touched", compare it with
	// errand boy's start time. If it was modified after, reload the config.
	if config == nil {
		return nil, errors.New("No loaded configuration.")
	}
	return config, nil
}
