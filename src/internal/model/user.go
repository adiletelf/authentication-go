package model

import (
	"github.com/google/uuid"
)

type User struct {
	UUID     uuid.UUID `json:"uuid" bson:"_id"`
	Password string    `json:"password"`
}

type UserRepo interface {
	Save(*User) (string, error)
	LoginCheck(uuid uuid.UUID, password string) (TokenDetails, error)
}
