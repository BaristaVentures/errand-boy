package repos

import (
	"encoding/json"
	"net/http"
)

func pullRequestHandler(res http.ResponseWriter, req *http.Request) {
	var prPayload PullRequest
	json.NewDecoder(req.Body).Decode(&prPayload)
	// // TODO: handle possible publisher errors.
	_ = eventsSubs["pr"].Publish(prPayload)
	res.WriteHeader(http.StatusOK)
}
