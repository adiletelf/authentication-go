package handler

import (
	"net/http"

	"github.com/adiletelf/authentication-go/internal/model"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	ur model.UserRepo
}

func New(ur model.UserRepo) *Handler {
	return &Handler{
		ur: ur,
	}
}

func (h *Handler) HandleHome(c *gin.Context) {
	c.String(200, "ok")
}

type tokenRequestBody struct {
	RefreshToken string `json:"refreshToken"`
}

func (h *Handler) Refresh(c *gin.Context) {
	var input tokenRequestBody

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, "ok")
}
