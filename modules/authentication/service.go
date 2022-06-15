package authentication

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/andydevstic/boilerplate-backend/modules/user"
	"github.com/andydevstic/boilerplate-backend/shared/custom"
	"github.com/andydevstic/boilerplate-backend/shared/interfaces"
	"github.com/andydevstic/boilerplate-backend/shared/utils"
	"github.com/andydevstic/boilerplate-backend/shared/utils/jwt"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
)

type AuthService struct{}

type IAuthService interface {
	Register(context context.Context, db interfaces.DBTX, dto *RegisterUserDTO) (err error)
	Login(context context.Context, db interfaces.DBTX, dto *LoginDTO) (user user.User, jwtToken string, err error)
}

func NewService() *AuthService {
	return &AuthService{}
}

func IsUserWithEmailExist(context context.Context, db interfaces.DBTX, email string) (bool, error) {
	findUserByEmailQuery := "SELECT id FROM users WHERE email = $1;"

	var userId string

	selectSmt, err := db.PrepareContext(context, findUserByEmailQuery)
	if err != nil {
		return false, fmt.Errorf("prepare find user query: %w", err)
	}

	row := selectSmt.QueryRowContext(context, email)
	err = row.Scan(&userId)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}

		return false, fmt.Errorf("check user exist: %w", err)
	}

	return true, nil
}

func (service *AuthService) Register(context context.Context, db interfaces.DBTX, dto *RegisterUserDTO) (err error) {
	isUserExist, err := IsUserWithEmailExist(context, db, dto.Email)
	if err != nil {
		return err
	}

	if isUserExist {
		return custom.NewError(http.StatusConflict, errors.New("register: email already used"))
	}

	insertQuery := `
		INSERT INTO users (name, email, password) VALUES ($1, $2, $3);
	`

	smt, err := db.PrepareContext(context, insertQuery)

	if err != nil {
		return fmt.Errorf("prepare insert query: %w", err)
	}

	defer smt.Close()

	hashedPassword, err := utils.HashFromPassword([]byte(dto.Password))
	if err != nil {
		return fmt.Errorf("hash password: %w", err)
	}

	fmt.Println(dto)

	_, err = smt.ExecContext(context, dto.Name, dto.Email, hashedPassword)

	return
}

func (service *AuthService) Login(context context.Context, db interfaces.DBTX, dto *LoginDTO) (user user.User, jwtToken string, err error) {
	selectQuery := `
		SELECT * FROM users WHERE email = $1;
	`

	smt, err := db.PrepareContext(context, selectQuery)
	if err != nil {
		return
	}

	defer smt.Close()

	row := smt.QueryRowContext(context, dto.Email)
	if err = row.Err(); err != nil {
		err = fmt.Errorf("query row: %w", err)

		return
	}

	err = row.Scan(&user.Id, &user.Email, &user.Name, &user.Type, &user.Status, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = custom.NewError(http.StatusUnauthorized, errors.New("username or password is incorrect"))

			return
		}

		err = fmt.Errorf("scan user fields: %w", err)

		return
	}

	isValid := utils.IsPasswordValid([]byte(user.Password), []byte(dto.Password))
	if !isValid {
		err = custom.NewError(http.StatusUnauthorized, fmt.Errorf("invalid username or password %w", err))

		return
	}

	payload := map[string]any{
		"id":     user.Id,
		"name":   user.Name,
		"email":  user.Email,
		"type":   user.Type,
		"status": user.Status,
	}

	jwtToken, err = jwt.SignPayload(payload)
	if err != nil {
		err = fmt.Errorf("sign payload: %w", err)
	}

	return user, jwtToken, err
}
