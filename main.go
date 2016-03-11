package main

import (
	"flag"
	"fmt"

	"github.com/BaristaVentures/errand-boy/server"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	port := flag.Int("p", 8080, "The port where Errand Boy will run.")
	configFilePath := flag.String("c", "./eb-conf.json", "The path to Errand Boy's config file.")
	flag.Parse()
	config, err := LoadConfig(*configFilePath)
	if err != nil {
		panic(err)
	}

	spew.Dump(config)
	s := server.Server{Port: *port}
	fmt.Printf("Errand Boy is ready to go on port %d.\n", s.Port)
	s.BootUp()
}
