package main

import (
	"flag"
	"github.com/BaristaVentures/errand-boy/config"
	"github.com/BaristaVentures/errand-boy/server"
	log "github.com/Sirupsen/logrus"
	"os"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	// use this to use the json format and it help the integration with 3rd parties log.SetFormatter(&log.JSONFormatter{})

	// we can also add hooks to send info to papertrail
	// package for papertrail hook https://github.com/polds/logrus-papertrail-hook

	// Output to stderr instead of stdout, could also be a file.
	log.SetOutput(os.Stderr)

	// Log info and higher we can change this to use warn and above on prod
	log.SetLevel(log.InfoLevel)
}

func main() {
	port := flag.Int("p", 8080, "The `port` where Errand Boy will run. Default: 8080")
	configFilePath := flag.String(
		"c",
		"./eb.conf.json",
		"The path to Errand Boy's `config file`. Default: ./eb.conf.json",
	)
	flag.Parse()

	_, err := config.Load(*configFilePath)
	if err != nil {
		panic(err)
	}

	s := server.Server{Port: *port}
	log.WithFields(log.Fields{
		"port": s.Port,
	}).Info("Errand Boy is ready to go on")
	s.BootUp()
}
