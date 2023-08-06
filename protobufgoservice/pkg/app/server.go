package app

import (
	api "ProtobufGoService/pkg/api"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router  *gin.Engine
	service api.ProtobufGoService
}

func NewServer(router *gin.Engine, service api.ProtobufGoService) *Server {
	return &Server{
		router:  router,
		service: service,
	}
}

func (s *Server) Run() error {
	// run the server through the router
	return s.SetupRouter().Run(":8080")
}
