package main

import (
	"footballanalyticshub/internal/db"
	"footballanalyticshub/internal/db/seed"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	database, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	err = seed.Run(database)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()
	router := gin.Default()
	router.GET("/health", returnMessage)
	router.Run("localhost:8080")
}

func returnMessage(c *gin.Context) {
	c.String(http.StatusOK, "Football Analytics Hub API is up")
}
