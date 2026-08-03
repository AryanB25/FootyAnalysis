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
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 20
	}
	min_rating, err := strconv.Atoi(c.Query("min_rating"))
	max_rating, err := strconv.Atoi(c.Query("max_rating"))
	players, err := h.repository.GetAllPlayers(page, limit, min_rating, max_rating)
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
