package main

import configuration "go-service-template/internal/config"

func main() {
	config, error := configuration.LoadConfig()
	if error != nil {
		panic(error)
	}

	panic(config)
}
