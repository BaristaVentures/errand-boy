package server

import (
	"net/http"
	"strconv"

	"github.com/BaristaVentures/errand-boy/routers/repos"
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

	baseRoute := r.PathPrefix("/")

	hooksSubRoute := r.PathPrefix("/hooks")
	hooksSubRouter := hooksSubRoute.Subrouter()

	reposSubRouter := hooksSubRouter.PathPrefix("/repos").Subrouter()

	baseRoute.Handler(repos.NormalizePRPayload(hooksSubRouter))

	repos.Route(reposSubRouter)
	http.ListenAndServe(":"+strconv.Itoa(s.Port), r)
}
