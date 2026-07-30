package main

import (
	"footballanalyticshub/internal/db"
	"footballanalyticshub/internal/db/seed"
	"footballanalyticshub/internal/players"
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
	repo := players.NewPlayerRepository(database)
	handler := players.NewPlayerHandlers(repo)
	router := gin.Default()
	router.GET("/health", returnMessage)
	router.GET("/players/search", handler.SearchPlayers)
	router.GET("/players/:id", handler.GetPlayerByID)
	router.GET("/players", handler.GetPlayers)
	router.Run("localhost:8080")
}

func returnMessage(c *gin.Context) {
	c.String(http.StatusOK, "Football Analytics Hub API is up")
}
