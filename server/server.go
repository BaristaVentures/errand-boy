package server

import (
	"net/http"
	"strconv"

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
	hooksRouter := r.PathPrefix("/hooks").Subrouter()
	// Add GitHub routes.
	github.Instance().SetUpRoutes(hooksRouter.PathPrefix("/gh").Subrouter())
	http.ListenAndServe(":"+strconv.Itoa(s.Port), r)
}
