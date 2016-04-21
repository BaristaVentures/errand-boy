package github

import (
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// Service specifies the methods a github service should implement.
type Service interface {
	CommentOnIssue(owner, repo string, number int, comment string) (*github.IssueComment, error)
}

type gitHubService struct {
	client *github.Client
}

func (service *gitHubService) CommentOnIssue(
	owner, repo string, number int, comment string) (*github.IssueComment, error) {

	issueComment := &github.IssueComment{Body: &comment}
	newComment, _, err := service.client.Issues.CreateComment(owner, repo, number, issueComment)
	return newComment, err
}

// New returns a new github.Service
func New(token string) Service {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	ghService := &gitHubService{client: github.NewClient(tc)}

	return ghService
}
