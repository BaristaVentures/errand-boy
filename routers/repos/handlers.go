package repos

import (
	"encoding/json"
	"net/http"
)

func pullRequestHandler(w http.ResponseWriter, r *http.Request) {
	var prPayload PullRequest
	json.NewDecoder(r.Body).Decode(&prPayload)
	// // TODO: handle possible publisher errors.
	_ = eventsSubs["pr"].Publish(prPayload)
	w.WriteHeader(http.StatusOK)
}
