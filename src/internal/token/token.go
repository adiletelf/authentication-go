package token

import (
	"strconv"
	"time"

	"github.com/adiletelf/authentication-go/internal/config"
	"github.com/adiletelf/authentication-go/internal/model"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func Generate(uuid uuid.UUID) (model.TokenDetails, error) {
	accessToken, err := generateAccessToken(uuid)
	if err != nil {
		return model.TokenDetails{}, err
	}
	refreshToken, err := generateRefreshToken(uuid)
	if err != nil {
		return model.TokenDetails{}, err
	}

	return model.TokenDetails{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func generateAccessToken(uuid uuid.UUID) (string, error) {
	cfg, _ := config.New()
	tokenLifespan, err := strconv.Atoi(cfg.AccessTokenMinuteLifespan)
	if err != nil {
		return "", err
	}
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["uuid"] = uuid
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(tokenLifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	signedToken, err := token.SignedString([]byte(cfg.ApiSecret))
	return signedToken, err
}

func generateRefreshToken(uuid uuid.UUID) (string, error) {
	cfg, _ := config.New()
	tokenLifespan, err := strconv.Atoi(cfg.RefreshTokenHourLifespan)
	if err != nil {
		return "", err
	}
	claims := jwt.MapClaims{}
	claims["uuid"] = uuid
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(tokenLifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	signedToken, err := token.SignedString([]byte(cfg.ApiSecret))
	return signedToken, err
}
