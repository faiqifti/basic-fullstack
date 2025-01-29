package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Message struct {
	Message string `json:"message"`
}

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, Message{Message: "Hello, World!"})
}

func main() {
	r := gin.Default()

	// Use CORS middleware
	r.Use(cors.Default())

	r.GET("/api/message", helloHandler)
	r.Run(":8080")
}
