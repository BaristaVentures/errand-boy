package main

import (
	"fmt"
	"strconv"

	"github.com/BaristaVentures/errand-boy/routers"
	"github.com/plimble/ace"
)

type server struct {
	port     int
	instance *ace.Ace
}

func (s *server) bootUp() {
	// Set the default ace instance with logging middleware.
	s.instance = ace.Default()
	// Initialize a new GitHub router.
	ghRouter := routers.NewGitHubRouter()
	// Add the new router's routes.
	s.setUpRoutes(ghRouter.GetRoutes())
	fmt.Printf("Errand Boy is listening for your commands on port %d.\n", s.port)
	s.instance.Run(":" + strconv.Itoa(s.port))
}

func (s *server) setUpRoutes(routes routers.Routes) {
	for _, r := range routes {
		s.instance.Handle(r.Method, r.Path, r.Handlers)
	}
}
