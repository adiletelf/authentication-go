package handler

import "github.com/adiletelf/authentication-go/internal/model"

type Handler struct {
	ur model.UserRepo
}

func New(ur model.UserRepo) *Handler {
	return &Handler{
		ur: ur,
	}
}
