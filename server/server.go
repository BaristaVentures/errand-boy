package server

import (
	"strconv"

	"github.com/BaristaVentures/errand-boy/routers/github"
	// Importing it as a blank package causes its init method to be called.
	_ "github.com/BaristaVentures/errand-boy/services/repotracker"
	"github.com/plimble/ace"
)

// Server represents the server's config.
type Server struct {
	Port int
}

// BootUp starts the server.
func (s *Server) BootUp() {
	// Set the default ace instance with logging middleware.
	a := ace.Default()
	hooksRouter := a.Router.Group("/hooks")
	// Add GitHub routes.
	github.Instance().SetUpRoutes(hooksRouter.Group("/gh"))
	a.Run(":" + strconv.Itoa(s.Port))
}
