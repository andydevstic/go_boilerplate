package pet

import (
	"context"
	"fmt"

	"github.com/andydevstic/boilerplate-backend/core"
	"github.com/andydevstic/boilerplate-backend/shared"
)

type PetService struct{}

type IPetService interface {
	shared.ICrudService[Pet]
}

func NewService() IPetService {
	return &PetService{}
}

func (*PetService) FindOne(ctx context.Context, criteria map[string]any) (Pet, error) {
	appState := core.GetAppState()

	foundPet := Pet{}

	tx := appState.Db.First(&foundPet, criteria)

	if err := tx.Error; err != nil {
		return foundPet, fmt.Errorf("find one user: %w", err)
	}

	return foundPet, nil
}

func (*PetService) Find(ctx context.Context, criteria map[string]any, limit, offset int) ([]Pet, error) {
	appState := core.GetAppState()

	foundPets := make([]Pet, 0, 10)

	tx := appState.Db.Model(&Pet{}).Find(foundPets).Limit(limit).Offset(offset)

	if err := tx.Error; err != nil {
		return foundPets, fmt.Errorf("find one pet: %w", err)
	}

	return foundPets, nil
}

func (*PetService) Create(ctx context.Context, payload map[string]any) error {
	appState := core.GetAppState()

	tx := appState.Db.Model(&Pet{}).Create(payload)
	if err := tx.Error; err != nil {
		return err
	}

	return nil
}

func (*PetService) Update(ctx context.Context, payload map[string]any) error {
	appState := core.GetAppState()

	tx := appState.Db.Model(&Pet{}).Updates(payload)
	if err := tx.Error; err != nil {
		return err
	}

	return nil
}

func (*PetService) Delete(ctx context.Context, criteria map[string]any) error {
	appState := core.GetAppState()

	tx := appState.Db.Model(&Pet{}).Delete(criteria)
	if err := tx.Error; err != nil {
		return err
	}

	return nil
}
