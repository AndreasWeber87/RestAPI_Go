package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type jsonMessage struct {
	Message string `json:"message"`
}

type jsonName struct {
	Name string `json:"name"`
}

func home(c *gin.Context) {
	var message = jsonMessage{
		Message: "Hello World! I'm the Go API.",
	}

	c.IndentedJSON(http.StatusOK, message)
}

func sayHelloGet(c *gin.Context) {
	name := c.Param("name")

	var message = jsonMessage{
		Message: "Hello " + name + "! I'm the Go API.",
	}

	c.IndentedJSON(http.StatusOK, message)
}

func sayHelloPost(c *gin.Context) {
	var name jsonName
	c.BindJSON(&name)

	var message = jsonMessage{
		Message: "Hello " + name.Name + "! I'm the Go API.",
	}

	c.IndentedJSON(http.StatusOK, message)
}
