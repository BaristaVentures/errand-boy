package config

import (
	"encoding/json"
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
	TrackerID int     `json:"tracker_id"`
	Repos     []*Repo `json:"repos"`
}

// Repo represents a repository
type Repo struct {
	Source string `json:"source"`
	Name   string `json:"name"`
	Token  string `json:"token"`
}

// Load  parses the config from a json file to a *Config and returns it.
func Load(path string) (*Config, error) {
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
	return config
}
