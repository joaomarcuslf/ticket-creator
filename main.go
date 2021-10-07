package main

import (
	"log"
	"os"

	_ "github.com/heroku/x/hmetrics/onload"
	appClient "github.com/joaomarcuslf/ticket-creator/clients/app"
	restClient "github.com/joaomarcuslf/ticket-creator/clients/rest"
)

func main() {
	port := os.Getenv("PORT")

	client := os.Getenv("CLIENT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	switch client {
	case "rest":
		rest := restClient.NewRestClient(port)

		rest.Initialize()
		break
	case "app":
	default:
		app := appClient.NewAppClient(port)

		app.Initialize()
		break
	}
}
