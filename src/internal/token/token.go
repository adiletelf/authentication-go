package token

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
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

func TokenValid(r *http.Request) error {
	cfg, _ := config.New()
	token := ExtractToken(r, "accessToken")
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(cfg.ApiSecret), nil
	})
	if err != nil {
		return err
	}
	return nil
}

func ExtractToken(r *http.Request, tokenName string) string {
	token := r.URL.Query().Get(tokenName)
	if token != "" {
		return token
	}
	bearerToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearerToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
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
