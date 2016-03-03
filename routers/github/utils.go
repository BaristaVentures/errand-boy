package github

import (
	"regexp"
	"strconv"
	"strings"
)

// The PR title should contain something like [PT projectID storyID]
var exp, _ = regexp.Compile("\\[PT\\s(\\d)+\\s(\\d)+\\]")

func parseTrackerCode(prTitle string) (int, int, bool) {
	codeSubStr := exp.FindString(prTitle)
	if len(codeSubStr) == 0 {
		return 0, 0, false
	}
	codeParts := strings.Split(strings.Trim(codeSubStr, "[]"), " ")
	if len(codeParts) != 3 {
		return 0, 0, false
	}
	projectID, err := strconv.Atoi(codeParts[1])
	if err != nil {
		return 0, 0, false
	}
	ID, err := strconv.Atoi(codeParts[2])
	return projectID, ID, false
}
