package jwt

import (
	"github.com/andydevstic/boilerplate-backend/config"
	"github.com/golang-jwt/jwt"
)

func SignPayload(payload jwt.MapClaims) (string, error) {
	config, err := config.GetConfig(".")
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	return token.SignedString([]byte(config.JwtSecret))
}
