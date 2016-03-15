package tracker

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/BaristaVentures/errand-boy/config"
	"github.com/BaristaVentures/errand-boy/routers"
	ghservice "github.com/BaristaVentures/errand-boy/services/github"
	"github.com/gorilla/mux"
	"github.com/salsita/go-pivotaltracker/v5/pivotal"
)

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

// GitHubPRURLData parses a pull request url and extracts its data.
func GitHubPRURLData(rawURL string) (owner, repo string, number int, err error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return owner, repo, number, err
	}

	splitPath := strings.Split(parsedURL.Path, "/")
	// After splitting it, the resulting array is [" ", <organization>, <repo>, "pull", <issue no.>]
	owner = splitPath[1]
	repo = splitPath[2]
	number, _ = strconv.Atoi(splitPath[4])
	return owner, repo, number, nil
}

func activityHandler(w http.ResponseWriter, r *http.Request) {
	activity := &ActivityPayload{}
	json.NewDecoder(r.Body).Decode(&activity)

	resource := activity.PrimaryResources[0]
	if resource.Kind != "story" {
		// Not a story, so return.
		w.WriteHeader(200)
		return
	}

	trackerProjectID := activity.Project.ID
	trackerStoryID := resource.ID

	trackerClient := pivotal.NewClient(config.Current().TrackerAPIToken)
	comments, _, err := trackerClient.Stories.ListComments(trackerProjectID, trackerStoryID)
	if err != nil {
		// log the error.
		fmt.Println(err.Error())
		w.WriteHeader(500)
		return
	}
	if len(comments) == 0 {
		// Comment on story saying the story doesn't follow errand boy's conventions.
		return
	}

	visitedRepos := make(map[string]bool)

	for _, comment := range comments {
		commentText := strings.Replace(comment.Text, " ", "", -1)
		// The errand boy-generated comment has the format "Check out the PR @ <pull request url>",
		// so the url should be second in splitText.
		splitText := strings.Split(commentText, "@")
		if len(splitText) < 2 {
			// Maybe it's a user comment, not necessarily a cause of error.
			continue
		}

		owner, repoName, issueNo, err := GitHubPRURLData(splitText[1])
		if err != nil || visitedRepos[repoName] {
			// If the right side of the @ wasn't a URL (maybe it was a @mention or an email),
			// or if we already processed a repo with this name, continue.
			continue
		}

		visitedRepos[repoName] = true

		projects := config.Current().Projects
		for _, project := range projects {
			if project.TrackerID == trackerProjectID {
				ghToken := project.Repos[repoName].Token
				comment := fmt.Sprintf("The [story](%s) related to this PR was %s", resource.URL, activity.Highlight)
				ghservice.New(ghToken).CommentOnIssue(owner, repoName, issueNo, comment)
				break
			}
		}
	}

	if len(visitedRepos) == 0 {
		// No comments were errand boy-compliant.
		// Comment on story saying the story doesn't follow errand boy's conventions.
	}
	w.WriteHeader(http.StatusOK)
}
