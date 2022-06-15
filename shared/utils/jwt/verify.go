package jwt

import (
	"fmt"

	"github.com/andydevstic/boilerplate-backend/config"
	"github.com/golang-jwt/jwt"
)

func VerifyToken(token string) (interface{}, error) {
	config, err := config.GetConfig(".")
	if err != nil {
		return "", err
	}

	parsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.JwtSecret), nil
	})

	if claims, ok := parsed.Claims.(jwt.MapClaims); ok && parsed.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("token is invalid: %v", err)
}
