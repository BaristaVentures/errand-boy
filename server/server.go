package server

import (
	"fmt"
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
	// Set the default ace instance with logging middleware.
	a := ace.Default()
	// Add GitHub routes.
	github.Instance().SetUpRoutes(a.Router)
	fmt.Printf("Errand Boy is listening for your commands on port %d.\n", s.Port)
	a.Run(":" + strconv.Itoa(s.Port))
}
