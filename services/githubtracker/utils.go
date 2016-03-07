package githubtracker

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// The PR title should contain something like [PT projectID storyID]
var exp, _ = regexp.Compile("\\[PT\\s(\\d)+\\s(\\d)+\\]")

func parseTrackerCode(prTitle string) (int, int, error) {
	codeSubStr := exp.FindString(prTitle)
	if len(codeSubStr) == 0 {
		return 0, 0, errors.New("Code format wasn't present.")
	}
	codeParts := strings.Split(strings.Trim(codeSubStr, "[]"), " ")
	if len(codeParts) != 3 {
		return 0, 0, errors.New("Code format wasn't valid.")
	}
	projectID, err := strconv.Atoi(codeParts[1])
	if err != nil {
		return 0, 0, errors.New("Code wasn't numeric.")
	}
	ID, err := strconv.Atoi(codeParts[2])
	return projectID, ID, nil
}
