package repos

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func replaceRequestBody(v interface{}, r *http.Request) {
	newBody, _ := json.Marshal(v)
	r.Body = ioutil.NopCloser(bytes.NewBuffer(newBody))
}

// NormalizePRPayload turns a bitbucket-specific PR payload into a general one.
func NormalizePRPayload(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pr := new(PullRequest)
		switch {
		case r.Header.Get("X-GitHub-Event") == "pull_request":
			// The request comes from GitHub.
			pr.HydrateFromGitHub(*r)
			replaceRequestBody(pr, r)
		case len(r.Header.Get("X-Event-Key")) > 0:
			// If the X-Event-Key header is set, It's bitbucket.
			pr.HydrateFromBitBucket(*r)
			replaceRequestBody(pr, r)
		}
		next.ServeHTTP(w, r)
	})
}
