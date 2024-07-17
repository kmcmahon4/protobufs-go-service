package app

import (
	"net/http"

	models "ProtobufGoService/pkg/models"

	"github.com/gin-gonic/gin"
)

func (s *Server) healthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/protobuf")

		// Read the Protobuf request data from the request body
		var request models.MyMessage
		if err := c.ShouldBind(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid protobuf data"})
			return
		}

		// Create a response Protobuf message (optional, if you want to respond with data)
		response := &models.Response{
			Status: "OK\n",
		}

		// Send the Protobuf response to the client
		c.ProtoBuf(http.StatusOK, response)

	}
}
