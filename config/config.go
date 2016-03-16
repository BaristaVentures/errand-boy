package config

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
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
	Source string `json:"source"`
	Token  string `json:"token"`
}

// Load  parses the config from a json file to a *Config and returns it.
func Load(path string) (*Config, error) {
	log.WithFields(log.Fields{
		"path": path,
	}).Info("Reading errand boy config")
	reader, err := os.Open(path)

	if err != nil {
		return nil, err
	}
	defer reader.Close()

	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	config = &Config{}
	json.Unmarshal(bytes, config)

	return config, nil
}

// Current returns the current config.
func Current() *Config {
	// TODO: Maybe we should check the last time the config file was "touched", compare it with
	// errand boy's start time. If it was modified after, reload the config.
	return config
}
