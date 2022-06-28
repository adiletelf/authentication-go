package model

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"id" bson:"_id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
}

type UserRepo interface {
	Save(*User) (string, error)
}
