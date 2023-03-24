package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type jsonMessage struct {
	Message string `json:"message"`
}

func home(c *gin.Context) {
	var response = jsonMessage{
		Message: "Hello World! I'm the Go API.",
	}

	c.IndentedJSON(http.StatusOK, response)
}
