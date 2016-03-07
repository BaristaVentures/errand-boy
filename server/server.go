package server

import (
	"strconv"

	ghRouter "github.com/BaristaVentures/errand-boy/routers/github"
	// Importing it as a blank package causes its init method to be called.
	_ "github.com/BaristaVentures/errand-boy/services/github"
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
	ghRouter.Instance().SetUpRoutes(a.Router)
	a.Run(":" + strconv.Itoa(s.Port))
}
