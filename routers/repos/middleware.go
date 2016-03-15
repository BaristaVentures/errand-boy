package repos

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// NormalizePRPayload turns a bitbucket-specific PR payload into a general one.
func NormalizePRPayload(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Header.Get("X-GitHub-Event") == "pull_request":
			// The request comes from GitHub.
			prPayload := &gitHubPRPayload{}
			json.NewDecoder(r.Body).Decode(&prPayload)
			genPayload := prPayload.ToGenericPR()
			genPayloadBytes, _ := json.Marshal(genPayload)
			r.Body = ioutil.NopCloser(bytes.NewBuffer(genPayloadBytes))
		case len(r.Header.Get("X-Event-Key")) > 0:
			// If the X-Event-Key header is set, It's bitbucket.
			prPayload := &bitBucketPRPayload{}
			json.NewDecoder(r.Body).Decode(&prPayload)
			genPayload := prPayload.ToGenericPR()
			genPayloadBytes, _ := json.Marshal(genPayload)
			r.Body = ioutil.NopCloser(bytes.NewBuffer(genPayloadBytes))
		}
		next.ServeHTTP(w, r)
	})
}
