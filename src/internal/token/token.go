package token

import (
	"github.com/adiletelf/authentication-go/internal/model"
	"github.com/google/uuid"
)

func Generate(uuid uuid.UUID) (model.TokenDetails, error) {
	return model.TokenDetails{
		AccessToken:  "access token stub",
		RefreshToken: "refresh token stub",
	}, nil
}
