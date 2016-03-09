package server

import (
	"net/http"
	"strconv"

	"github.com/BaristaVentures/errand-boy/routers/bitbucket"
	"github.com/BaristaVentures/errand-boy/routers/github"
	// Importing it as a blank package causes its init method to be called.
	_ "github.com/BaristaVentures/errand-boy/services/repotracker"
	"github.com/gorilla/mux"
)

// Server represents the server's config.
type Server struct {
	Port int
}

// BootUp starts the server.
func (s *Server) BootUp() {
	r := mux.NewRouter()
	r.StrictSlash(true)
	hooksSubRouter := r.PathPrefix("/hooks").Subrouter()
	// Add GitHub routes.
	hooksSubRouter.PathPrefix("/gh").Handler(github.Router())
	// Add BitBucket routes.
	hooksSubRouter.PathPrefix("/bb").Handler(bitbucket.NormalizePRPayload(bitbucket.Router()))
	// Start listening.
	http.ListenAndServe(":"+strconv.Itoa(s.Port), r)
}
