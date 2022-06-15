package user

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andydevstic/boilerplate-backend/shared/constants"
	"github.com/andydevstic/boilerplate-backend/shared/interfaces"
	"github.com/andydevstic/boilerplate-backend/shared/utils"
	"github.com/andydevstic/boilerplate-backend/shared/utils/filter"
	"github.com/doug-martin/goqu/v9"
	"github.com/rs/zerolog/log"
)

type UserService struct{}

type IUserService interface {
	FindUserByEmail(context context.Context, db interfaces.DBTX, email string) (user User, err error)
	FindUserById(context context.Context, db interfaces.DBTX, userId int) (user User, err error)
	CreateUserAdmin(context context.Context, db interfaces.DBTX, dto *UpsertUserAdminDTO) error
	UpdateUserById(context context.Context, db interfaces.DBTX, userId int, dto *UpdateUserDTO) error
	UpdateUserByIdAdmin(context context.Context, db interfaces.DBTX, userId int, dto *UpsertUserAdminDTO) error
	FindUsersAdmin(context context.Context, db interfaces.DBTX, dto *FindUsersAdminDTO) ([]User, error)
}

func NewService() *UserService {
	return &UserService{}
}

func (*UserService) FindUserByEmail(context context.Context, db interfaces.DBTX, email string) (user User, err error) {
	findByEmailQuery := `
		SELECT id, name, email, status, type, password
		FROM users
		WHERE email = $1;
	`

	smt, err := db.PrepareContext(context, findByEmailQuery)

	if err != nil {
		return
	}

	defer smt.Close()

	row := smt.QueryRow(sql.Named("email", email))
	if err = row.Err(); err != nil {
		return
	}

	err = row.Scan(&user.Id, &user.Name, &user.Email, &user.Status, &user.Type, &user.Password)

	return
}

func (*UserService) FindUserById(context context.Context, db interfaces.DBTX, userId int) (user User, err error) {
	findByIdQuery := `
		SELECT id, name, email, status, type
		FROM users
		WHERE id = $1;
	`

	smt, err := db.PrepareContext(context, findByIdQuery)

	if err != nil {
		return
	}

	defer smt.Close()

	row := smt.QueryRow(sql.Named("id", userId))
	if err = row.Err(); err != nil {
		return
	}

	err = row.Scan(&user.Id, &user.Name, &user.Email, &user.Status, &user.Type)

	return
}

func (*UserService) CreateUserAdmin(context context.Context, db interfaces.DBTX, dto *UpsertUserAdminDTO) error {
	insertQuery := `
		INSERT INTO users (name, email, type, status, password) VALUES ($1, $2, $3, $4, $5);
	`

	smt, err := db.PrepareContext(context, insertQuery)

	if err != nil {
		log.Error().Msg(fmt.Sprintf("Failed to prepare create user query: %s", err.Error()))

		return err
	}

	defer smt.Close()

	hashedPassword, err := utils.HashFromPassword([]byte(dto.Password))
	if err != nil {
		log.Error().Msg("Failed to has user password")

		return err
	}

	_, err = smt.ExecContext(context, dto.Name, dto.Email, dto.Type, dto.Status, hashedPassword)

	if err != nil {
		return err
	}

	return nil
}

func (*UserService) UpdateUserById(context context.Context, db interfaces.DBTX, userId int, dto *UpdateUserDTO) error {
	updateQuery := `
		UPDATE users SET name = $1;
	`

	smt, err := db.PrepareContext(context, updateQuery)
	if err != nil {
		log.Error().Msg(err.Error())

		return err
	}

	defer smt.Close()

	_, err = smt.ExecContext(context, dto.Name)
	if err != nil {
		return err
	}

	return nil
}

func (*UserService) UpdateUserByIdAdmin(context context.Context, db interfaces.DBTX, userId int, dto *UpsertUserAdminDTO) error {
	updateQuery := `
		UPDATE users SET name = $1, email = $2, type = $3, status = $4 WHERE id = $5;
	`

	smt, err := db.PrepareContext(context, updateQuery)
	if err != nil {
		log.Error().Msg(err.Error())

		return err
	}

	defer smt.Close()

	_, err = smt.ExecContext(context, dto.Name, dto.Email, dto.Type, dto.Status, userId)

	if err != nil {
		return err
	}

	return nil
}

func (*UserService) FindUsersAdmin(context context.Context, db interfaces.DBTX, dto *FindUsersAdminDTO) ([]User, error) {
	selectStatement := goqu.Dialect(constants.DbDialect).From("users").Limit(dto.Limit).Offset(dto.Offset)

	if dto.Email != "" {
		emailFilter, err := filter.ParseFilterString("email", dto.Email)
		if err != nil {
			return []User{}, err
		}

		selectStatement.Where(emailFilter)
	}

	if dto.Name != "" {
		nameFilter, err := filter.ParseFilterString("name", dto.Name)
		if err != nil {
			return []User{}, err
		}

		selectStatement.Where(nameFilter)
	}

	if dto.Type != "" {
		typeFilter, err := filter.ParseFilterString("type", dto.Type)
		if err != nil {
			return []User{}, err
		}

		selectStatement.Where(typeFilter)
	}

	if dto.Status != "" {
		statusFilter, err := filter.ParseFilterString("status", dto.Status)
		if err != nil {
			return []User{}, err
		}

		selectStatement.Where(statusFilter)
	}

	query, _, err := selectStatement.ToSQL()
	if err != nil {
		return []User{}, err
	}

	rows, err := db.QueryContext(context, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]User, 0, 20)

	for rows.Next() {
		user := User{}

		err = rows.Scan(
			&user.Id,
			&user.Email,
			&user.Name,
			&user.Type,
			&user.Status,
			&user.Password,
		)

		if err != nil {
			return nil, err
		}

		result = append(result, user)
	}

	return result, nil
}
