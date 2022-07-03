package authentication

import (
	"fmt"

	"github.com/andydevstic/boilerplate-backend/modules/user"
	jwtutils "github.com/andydevstic/boilerplate-backend/shared/utils/jwt"
	"github.com/golang-jwt/jwt"
	"github.com/mitchellh/mapstructure"
)

type AuthService struct{}

type IAuthService interface {
	ValidateUserPassword(password, userPassword string) error
	GenerateJwtFromUser(*user.User) (string, error)
}

func NewService() IAuthService {
	return &AuthService{}
}

func (*AuthService) ValidateUserPassword(password, userPassword string) error {
	return nil
}

func (*AuthService) GenerateJwtFromUser(user *user.User) (string, error) {
	jwtClaim := jwt.MapClaims{}

	err := mapstructure.Decode(user, jwtClaim)
	if err != nil {
		err = fmt.Errorf("decode user into jwt: %w", err)

		return "", err
	}

	jwtToken, err := jwtutils.SignPayload(jwtClaim)
	if err != nil {
		err = fmt.Errorf("sign jwt payload: %w", err)

		return "", err
	}

	return jwtToken, nil
}
