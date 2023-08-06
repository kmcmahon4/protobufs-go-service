package main

import (
	pkg "ProtobufGoService/pkg/api"
	"ProtobufGoService/pkg/app"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}

func run() error {
	// create router dependency
	router := gin.Default()
	router.Use(cors.Default())

	// create service
	service := pkg.NewProtobufGoService()

	server := app.NewServer(router, service)

	// start the server
	err := server.Run()
	if err != nil {
		return err
	}

	return nil

}
