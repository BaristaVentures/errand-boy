package main

import "github.com/BaristaVentures/errand-boy/server"

func main() {
	s := server.Server{Port: 8080}
	s.BootUp()
}
