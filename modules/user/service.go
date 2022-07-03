package user

import (
	"context"
	"fmt"

	"github.com/andydevstic/boilerplate-backend/core"
	"github.com/andydevstic/boilerplate-backend/shared"
)

type UserService struct{}

type IUserService interface {
	shared.ICrudService[User]
}

func NewService() IUserService {
	return &UserService{}
}

func (*UserService) FindOne(ctx context.Context, criteria map[string]any) (User, error) {
	appState := core.GetAppState()

	foundUser := User{}

	tx := appState.Db.Model(&User{}).First(criteria)

	if err := tx.Error; err != nil {
		return foundUser, fmt.Errorf("find one user: %w", err)
	}

	return foundUser, nil
}

func (*UserService) Find(ctx context.Context, criteria map[string]any, limit, offset int) ([]User, error) {
	appState := core.GetAppState()

	foundUsers := make([]User, 0, 10)

	tx := appState.Db.Model(&User{}).Find(foundUsers).Limit(limit).Offset(offset)

	if err := tx.Error; err != nil {
		return foundUsers, fmt.Errorf("find one user: %w", err)
	}

	return foundUsers, nil
}

func (*UserService) Create(ctx context.Context, payload map[string]any) error {
	appState := core.GetAppState()

	tx := appState.Db.Model(&User{}).Create(payload)
	if err := tx.Error; err != nil {
		return err
	}

	return nil
}

func (*UserService) Update(ctx context.Context, payload map[string]any) error {
	appState := core.GetAppState()

	tx := appState.Db.Model(&User{}).Updates(payload)
	if err := tx.Error; err != nil {
		return err
	}

	return nil
}

func (*UserService) Delete(ctx context.Context, criteria map[string]any) error {
	appState := core.GetAppState()

	tx := appState.Db.Model(&User{}).Delete(criteria)
	if err := tx.Error; err != nil {
		return err
	}

	return nil
}
