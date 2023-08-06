package app

import (
	"github.com/gin-gonic/gin"
)

func (s *Server) SetupRouter() *gin.Engine {
	router := s.router
	router.GET("/healthCheck", s.healthCheck())

	return router
}
