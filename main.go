package main

import (
	"flag"
	"fmt"

	"github.com/BaristaVentures/errand-boy/config"
	"github.com/BaristaVentures/errand-boy/server"
)

func main() {
	port := flag.Int("p", 8080, "The `port` where Errand Boy will run. Default: 8080")
	configFilePath := flag.String(
		"c",
		"./eb-conf.json",
		"The path to Errand Boy's `config file`. Default: ./eb.conf.json",
	)
	flag.Parse()

	_, err := config.Load(*configFilePath)
	if err != nil {
		panic(err)
	}

	s := server.Server{Port: *port}
	fmt.Printf("Errand Boy is ready to go on port %d.\n", s.Port)
	s.BootUp()
}
