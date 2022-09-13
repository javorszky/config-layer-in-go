package main

import (
	"log"

	"github.com/javorszky/config-layer/config"
	"github.com/javorszky/config-layer/service"
)

func main() {
	// the config module whose entire reason to exist it to parse environment
	// variables and make sure they exist, and are in a good set of values.
	//
	// @see config/config.go file.
	appConfig, err := config.Parse()
	if err != nil {
		// something went wrong, stop the world
		log.Fatalf("config parsing failed: %s", err.Error())
	}

	// pass in the parsed config struct, so app doesn't need to deal
	// with reading env vars. By this time, it's already been done.
	// Every time config is passed, you know it has every information
	// already ready, and in a format that other parts would need.
	//
	// @see service/service.go file.
	app := service.New(appConfig)
	app.Start()
}
