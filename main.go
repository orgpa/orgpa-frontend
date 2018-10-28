package main

import (
	"log"
	"orgpa-frontend/configuration"
	"orgpa-frontend/orgpa"
)

func main() {
	config, _ := configuration.ExtractConfiguration("configuration.json")

	template := config.InitTemplate()

	err := orgpa.Run(config, template)
	if err != nil {
		log.Fatal(err.Error())
	}
}
