package main

import (
	pkg "ProtobufGoService/pkg/api"
	"ProtobufGoService/pkg/app"
	"fmt"
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
	router := gin.Default()
	router.Use(cors.Default())

	service := pkg.NewProtobufGoService()
	if router == nil {
		fmt.Println("router nil")
	}

	server := app.NewServer(router, service)

	// start the server
	err := server.Run()
	if err != nil {
		return err
	}

	return nil

}
