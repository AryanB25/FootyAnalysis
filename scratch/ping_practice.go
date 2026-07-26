package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/ping", returnPong)
	router.Run("localhost:8080")
}

func returnPong(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "pong")
}