package bash

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/BaristaVentures/errand-boy/config"
	"github.com/hooklift/assert"
)

func TestCreateCommandsFile(t *testing.T) {
	name := "awesome-repo"
	filename := "./" + name + "-commands.sh"
	repo := &config.Repo{
		Commands: []string{"ls", "cd ~", "echo AW YIS"},
	}

	expFileContents := fmt.Sprintf("%s\n%s\n%s\n", repo.Commands[0], repo.Commands[1], repo.Commands[2])

	err := createCommandsFile(name, repo, filename)
	assert.Ok(t, err)
	file, err := os.Open(filename)
	assert.Ok(t, err)
	defer file.Close()
	defer os.Remove(filename)
	fileBytes, err := ioutil.ReadAll(file)
	assert.Ok(t, err)
	actFileContents := string(fileBytes)
	assert.Equals(t, expFileContents, actFileContents)
}

func TestCreateCommandsFileNoCommands(t *testing.T) {
	name := "awesome-repo"
	filename := "./" + name + "-commands.sh"
	repo := &config.Repo{
		Commands: []string{},
	}

	err := createCommandsFile(name, repo, filename)
	assert.Cond(t, err != nil, "err shouldn't be nil when there's no commands.")
}
