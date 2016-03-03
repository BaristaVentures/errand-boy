package main

import (
	"fmt"

	"github.com/BaristaVentures/errand-boy/server"
)

func main() {
	s := server.Server{Port: 8080}
	fmt.Printf("Errand Boy is listening for your commands on port %d.\n", s.Port)
	s.BootUp()
}
