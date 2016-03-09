package repos

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// NormalizePRPayload turns a bitbucket-specific PR payload into a general one.
func NormalizePRPayload(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url, _ := url.Parse(r.URL.Path)
		fmt.Println(url.Host)
		// NOTICE: url.Host == "" when on localhost.
		switch url.Host {
		case "github.com":
			prPayload := &GitHubPRPayload{}
			json.NewDecoder(r.Body).Decode(&prPayload)
			genPayload := prPayload.ToGeneric()
			genPayloadBytes, _ := json.Marshal(genPayload)
			r.Body = ioutil.NopCloser(bytes.NewBuffer(genPayloadBytes))
			next.ServeHTTP(w, r)
		case "bitbucket.com":
			prPayload := &BitBucketPRPayload{}
			json.NewDecoder(r.Body).Decode(&prPayload)
			genPayload := prPayload.ToGeneric()
			genPayloadBytes, _ := json.Marshal(genPayload)
			r.Body = ioutil.NopCloser(bytes.NewBuffer(genPayloadBytes))
			next.ServeHTTP(w, r)
		default:
			next.ServeHTTP(w, r)
		}
	})
}
