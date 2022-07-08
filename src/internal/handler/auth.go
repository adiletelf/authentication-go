package handler

import (
	"net/http"

	"github.com/adiletelf/authentication-go/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RegisterBody struct {
	Password string `json:"password"`
}

func (h *Handler) Register(c *gin.Context) {
	var input RegisterBody

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := model.User{
		UUID:     uuid.New(),
		Password: input.Password,
	}

	insertedID, err := h.ur.Save(&u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "uuid": insertedID})
}

type LoginBody struct {
	UUID     uuid.UUID `json:"uuid"`
	Password string    `json:"password"`
}

func (h *Handler) Login(c *gin.Context) {
	var input LoginBody

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := model.User{
		UUID:     input.UUID,
		Password: input.Password,
	}

	tokenDetails, err := h.ur.LoginCheck(u.UUID, u.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"accessToken":  tokenDetails.AccessToken,
		"refreshToken": tokenDetails.RefreshToken,
	})
}
