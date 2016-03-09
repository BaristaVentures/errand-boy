package server

import (
	"net/http"
	"strconv"

	// Importing it as a blank package causes its init method to be called.
	"github.com/BaristaVentures/errand-boy/routers/repos"
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

	baseRoute := r.PathPrefix("/")

	hooksSubRoute := r.PathPrefix("/hooks")
	hooksSubRouter := hooksSubRoute.Subrouter()

	reposSubRouter := hooksSubRouter.PathPrefix("/repos").Subrouter()
	// bbSubRouter := hooksSubRouter.PathPrefix("/bb").Subrouter()
	// ghSubRouter := hooksSubRouter.PathPrefix("/gh").Subrouter()

	baseRoute.Handler(repos.NormalizePRPayload(hooksSubRouter))

	repos.Router(reposSubRouter)
	// Add GitHub routes.
	// github.Router(ghSubRouter)
	// Add BitBucket routes.
	// bitbucket.Router(bbSubRouter)
	// Start listening.
	http.ListenAndServe(":"+strconv.Itoa(s.Port), r)
}
