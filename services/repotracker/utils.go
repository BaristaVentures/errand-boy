package repotracker

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/BaristaVentures/errand-boy/config"
	"github.com/BaristaVentures/errand-boy/routers/repos"
)

// The PR title should contain something like [storyID]
var exp, _ = regexp.Compile("\\[(\\d)+\\]")

func GetTrackerData(pr *repos.PullRequest) (projectID, storyID int, err error) {
	codeSubStr := exp.FindString(pr.Title)
	if len(codeSubStr) == 0 {
		return 0, 0, errors.New("Code format wasn't present.")
	}
	// Error can be ignored because at this point, the regexp guarantees there's an int there.
	storyID, _ = strconv.Atoi(strings.Trim(codeSubStr, "[]"))
	projectID, ok := getProjectIDForRepo(pr.Repo)
	if !ok {
		return 0, 0, errors.New("No Project Tracker ID for repo " + pr.Repo + " in current config.")
	}
	return projectID, storyID, nil
}

func getProjectIDForRepo(repo string) (int, bool) {
	curConfig, err := config.Current()
	if err != nil {
		return 0, false
	}
	projects := curConfig.Projects
	for _, p := range projects {
		if _, ok := p.Repos[repo]; ok {
			return p.TrackerID, true
		}
	}
	return 0, false
}
