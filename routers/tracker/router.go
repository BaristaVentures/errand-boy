package tracker

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/BaristaVentures/errand-boy/config"
	"github.com/BaristaVentures/errand-boy/routers"
	ghservice "github.com/BaristaVentures/errand-boy/services/github"
	"github.com/BaristaVentures/errand-boy/services/tracker"
	"github.com/Sirupsen/logrus"
	"github.com/deckarep/golang-set"
	"github.com/gorilla/mux"
)

var supportedHighlights = mapset.NewSetFromSlice([]interface{}{"accepted", "rejected"})

type pullRequestData struct {
	Owner    string
	RepoName string
	Number   int
	Host     string
}

// Route returns a configured *mux.Router.
func Route(router *mux.Router) *mux.Router {
	routes := &routers.Routes{
		&routers.Route{
			Path:    "/activity",
			Method:  "POST",
			Handler: activityHandler,
		},
	}
	for _, r := range *routes {
		router.Methods(r.Method).Path(r.Path).Handler(r.Handler)
	}

	return router
}

func activityHandler(w http.ResponseWriter, r *http.Request) {
	activity := &ActivityPayload{}
	json.NewDecoder(r.Body).Decode(&activity)

	resource := activity.PrimaryResources[0]
	if resource.Kind != "story" || !supportedHighlights.Contains(activity.Highlight) {
		// Not a story, so return.
		w.WriteHeader(http.StatusOK)
		return
	}

	apiToken := os.Getenv(config.Current().TrackerAPIToken)

	trackerService := tracker.New(apiToken)
	comments, err := trackerService.GetStoryComments(activity.Project.ID, resource.ID)
	if err != nil {
		logrus.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if len(comments) == 0 {
		// Comment on story saying the story doesn't follow errand boy's conventions.
		comment := `This story doesn't have any comment.
    Errand Boy relies on story comments to map them to their repository.`
		trackerService.CommentOnStory(activity.Project.ID, resource.ID, comment)
		return
	}

	visitedRepos := make(map[string]bool)

	for _, comment := range comments {
		prData, err := extractDataFromComment(comment.Text)
		if err != nil || visitedRepos[prData.RepoName] {
			// If the right side of the "@" wasn't a valid URL (maybe it was a @mention or an email),
			// or if we already processed a repo with this name, continue.
			continue
		}

		visitedRepos[prData.RepoName] = true
		projects := config.Current().Projects
		for _, project := range projects {
			if project.TrackerID == activity.Project.ID {
				ghTokenEnvVar := project.Repos[prData.RepoName].Token
				ghToken := os.Getenv(ghTokenEnvVar)
				if len(ghToken) == 0 {
					logrus.Error("No value set for env var " + ghTokenEnvVar)
					break
				}
				fmtStr := "%s %s the [story](%s) related to this PR in Pivotal Tracker."
				comment := fmt.Sprintf(fmtStr, activity.Actor.Name, activity.Highlight, resource.URL)
				switch prData.Host {
				case "github.com":
					ghservice.New(ghToken).CommentOnIssue(prData.Owner, prData.RepoName, prData.Number, comment)
				}
				break
			}
		}
	}
	if len(visitedRepos) == 0 {
		// No comments were errand boy-compliant.
		comment := `This story doesn't have any comment that matches Errand Boy's conventions.
    Errand Boy relies on story comments to map them to their repository.`
		trackerService.CommentOnStory(activity.Project.ID, resource.ID, comment)
	}
	w.WriteHeader(http.StatusOK)
}

// GitHubPRURLData parses a pull request url and extracts its data.
func extractDataFromComment(comment string) (*pullRequestData, error) {
	// Remove all spaces to make it easier to process.
	commentText := strings.Replace(comment, " ", "", -1)
	// The errand boy-generated comment has the format "Check out the PR @ <pull request url>",
	// so the url should be second in splitText.
	splitText := strings.Split(commentText, "@")
	rawURL := splitText[1]
	if len(rawURL) == 0 {
		return nil, errors.New("No URL present or invalid format.")
	}
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	// url.Parse doesn't return an error for URLs like "0892734" or "lasd;asd", so we have to make
	// sure the host was set.
	if len(parsedURL.Host) == 0 {
		return nil, errors.New("Invalid URL: " + rawURL)
	}

	splitPath := strings.Split(parsedURL.Path, "/")
	// After splitting it, the resulting array is [" ", <organization>, <repo>, "pull", <issue no.>]
	owner := splitPath[1]
	repo := splitPath[2]
	number, _ := strconv.Atoi(splitPath[4])
	return &pullRequestData{owner, repo, number, parsedURL.Host}, nil
}
