package repos

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func genericBody(payload PRConverter, r *http.Request) {
	json.NewDecoder(r.Body).Decode(&payload)
	genPayload := payload.ToGenericPR()
	genPayloadBytes, _ := json.Marshal(genPayload)
	r.Body = ioutil.NopCloser(bytes.NewBuffer(genPayloadBytes))
}

// NormalizePRPayload turns a bitbucket-specific PR payload into a general one.
func NormalizePRPayload(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Header.Get("X-GitHub-Event") == "pull_request":
			// The request comes from GitHub.
			prPayloadStruct := &gitHubPRPayload{}
			genericBody(prPayloadStruct, r)
		case len(r.Header.Get("X-Event-Key")) > 0:
			// If the X-Event-Key header is set, It's bitbucket.
			prPayloadStruct := &bitBucketPRPayload{}
			genericBody(prPayloadStruct, r)
		}
		next.ServeHTTP(w, r)
	})
}
