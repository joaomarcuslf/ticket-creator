package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/heroku/x/hmetrics/onload"
	appClient "github.com/joaomarcuslf/ticket-creator/clients/app"
	grpcClient "github.com/joaomarcuslf/ticket-creator/clients/grpc"
	restClient "github.com/joaomarcuslf/ticket-creator/clients/rest"
)

func main() {
	port := os.Getenv("PORT")

	client := os.Getenv("CLIENT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	switch client {
	case "grpc":
		grpc := grpcClient.NewGrpcClient(port)

		grpc.Initialize()
		break
	case "rest":
		rest := restClient.NewRestClient(port)

		rest.Initialize()
		break
	case "app":
		app := appClient.NewAppClient(port)

		app.Initialize()
		break
	default:
		intVar, err := strconv.Atoi(port)

		if err != nil {
			log.Fatal(err)
		}

		grpc := grpcClient.NewGrpcClient(fmt.Sprintf("%v", intVar-1))
		rest := restClient.NewRestClient(fmt.Sprintf("%v", intVar+1))
		app := appClient.NewAppClient(fmt.Sprintf("%v", intVar))

		go func() { grpc.Initialize() }()
		fmt.Println("Initialize grpc on :" + fmt.Sprintf("%v", intVar-1))
		go func() { rest.Initialize() }()
		fmt.Println("Initialize rest on :" + fmt.Sprintf("%v", intVar+1))

		app.Initialize()
	}
}
