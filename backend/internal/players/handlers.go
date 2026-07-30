package players

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Handlers struct {
	repository *Repository
}

func NewPlayerHandlers(repository *Repository) *Handlers {
	return &Handlers{
		repository: repository,
	}
}

func (h *Handlers) GetPlayers(c *gin.Context) {
	players, err := h.repository.GetAllPlayers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, players)
}

func (h *Handlers) GetPlayerByID(c *gin.Context) {
	number, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	player, err := h.repository.GetPlayerByID(number)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, player)
}

func (h *Handlers) SearchPlayers(c *gin.Context) {
	players, err := h.repository.SearchPlayers(c.Query("name"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, players)
}
