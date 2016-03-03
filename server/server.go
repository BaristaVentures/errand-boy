package server

import (
	"strconv"

	"github.com/BaristaVentures/errand-boy/routers/github"
	"github.com/plimble/ace"
)

// Server represents the server's config.
type Server struct {
	Port int
}

// BootUp starts the server.
func (s *Server) BootUp() {
	// Set the default ace instance with logging mIDdleware.
	a := ace.Default()
	// Add GitHub routes.
	github.Instance().SetUpRoutes(a.Router)
	a.Run(":" + strconv.Itoa(s.Port))
}
