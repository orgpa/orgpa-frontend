package main

import (
	"log"
	"orgpa-frontend/configuration"
	"orgpa-frontend/orgpa"
)

func main() {
	config, err := configuration.ExtractConfiguration("configuration.json")
	if err != nil {
		log.Fatal(err.Error())
	}

	template := config.InitTemplate()

	err = orgpa.Run(config, template)
	if err != nil {
		log.Fatal(err.Error())
	}
}
