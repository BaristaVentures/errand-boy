package main

import (
	"flag"
	"fmt"

	"github.com/BaristaVentures/errand-boy/server"
)

func main() {
	port := flag.Int("p", 8080, "The port where Errand Boy will run.")
	flag.Parse()
	s := server.Server{Port: *port}
	fmt.Printf("Errand Boy is ready to go on port %d.\n", s.Port)
	s.BootUp()
}
