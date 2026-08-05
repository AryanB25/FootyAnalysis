package players

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
	if page <= 0 {
		ErrorResponse(http.StatusBadRequest, "page must be greater than 0", c)
		return
	}
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 20
	}
	if limit > 100 || limit <= 0 {
		ErrorResponse(http.StatusBadRequest, "limit must be between 1 and 100", c)
		return
	}
	minRating, err := strconv.Atoi(c.Query("min_rating"))
	if err != nil {
		minRating = 0
	}
	if minRating > 99 || minRating < 0 {
		ErrorResponse(http.StatusBadRequest, "min rating must be between 0 and 99", c)
		return
	}
	maxRating, err := strconv.Atoi(c.Query("max_rating"))
	if err != nil {
		maxRating = 0
	}
	if maxRating > 99 || maxRating < 0 {
		ErrorResponse(http.StatusBadRequest, "max rating must be between 0 and 99", c)
		return
	}
	if minRating > 0 && maxRating > 0 && minRating > maxRating {
		ErrorResponse(http.StatusBadRequest, "min rating cannot be greater than max rating", c)
		return
	}
	players, err := h.repository.GetAllPlayers(page, limit, minRating, maxRating)
	if err != nil {
		ErrorResponse(http.StatusInternalServerError, err.Error(), c)
		return
	}
	c.JSON(http.StatusOK, players)
}

func (h *Handlers) GetPlayerByID(c *gin.Context) {
	number, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ErrorResponse(http.StatusBadRequest, err.Error(), c)
		return
	}
	player, err := h.repository.GetPlayerByID(number)
	if errors.Is(err, sql.ErrNoRows) {
		ErrorResponse(http.StatusNotFound, "player not found", c)
		return
	}
	if err != nil {
		ErrorResponse(http.StatusInternalServerError, err.Error(), c)
		return
	}
	c.JSON(http.StatusOK, player)
}

func (h *Handlers) SearchPlayers(c *gin.Context) {
	players, err := h.repository.SearchPlayers(c.Query("name"))
	if err != nil {
		ErrorResponse(http.StatusInternalServerError, err.Error(), c)
		return
	}
	c.JSON(http.StatusOK, players)
}

func (h *Handlers) ComparePlayers(c *gin.Context) {
	idPlayer1, err := strconv.Atoi(c.Query("id1"))
	if err != nil {
		ErrorResponse(http.StatusBadRequest, err.Error(), c)
		return
	}
	player1, err := h.repository.GetPlayerByID(idPlayer1)
	if errors.Is(err, sql.ErrNoRows) {
		ErrorResponse(http.StatusNotFound, "player not found", c)
		return
	}
	if err != nil {
		ErrorResponse(http.StatusInternalServerError, err.Error(), c)
		return
	}
	idPlayer2, err := strconv.Atoi(c.Query("id2"))
	if err != nil {
		ErrorResponse(http.StatusBadRequest, err.Error(), c)
		return
	}
	player2, err := h.repository.GetPlayerByID(idPlayer2)
	if errors.Is(err, sql.ErrNoRows) {
		ErrorResponse(http.StatusNotFound, "player not found", c)
		return
	}
	if err != nil {
		ErrorResponse(http.StatusInternalServerError, err.Error(), c)
		return
	}

	convertedPlayer1 := convertToComparisonStats(player1)
	convertedPlayer2 := convertToComparisonStats(player2)

	c.JSON(http.StatusOK, PlayerComparison{Player1: convertedPlayer1, Player2: convertedPlayer2})
}

func ErrorResponse(status int, message string, c *gin.Context) {
	c.JSON(status, gin.H{"error": message})
}
