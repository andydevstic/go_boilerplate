package pet

import (
	"fmt"

	"github.com/andydevstic/boilerplate-backend/shared"
	"github.com/andydevstic/boilerplate-backend/shared/utils/filter"
	"gorm.io/gorm"
)

type Pet struct {
	gorm.Model
	Name            string `gorm:"column:name"`
	FoundAddress    string `gorm:"column:foundAddress"`
	LastSeenAddress string `gorm:"column:lastSeenAddress"`
	PetTypeID       uint   `gorm:"column:typeID"`
	PetSpecieID     uint   `gorm:"column:specieID"`
	PetBreedID      uint   `gorm:"column:breedID"`
	PetFurLengthID  uint   `gorm:"column:furLengthID"`
	PetFurColorID   uint   `gorm:"column:furColorID"`
	PetSizeID       uint   `gorm:"column:sizeID"`
	PetAgeID        uint   `gorm:"column:ageID"`
	PetType         PetType
	PetSpecie       PetSpecie
	PetBreed        PetBreed
	PetFurLength    PetFurLength
	PetFurColor     PetFurColor
	PetSize         PetSize
	PetAge          PetAge
}

type PetType struct {
	gorm.Model
	Name string `json:"name"`
}

type PetSpecie struct {
	gorm.Model
	Name string `json:"name"`
}

type PetBreed struct {
	gorm.Model
	Name string `json:"name"`
}

type PetFurLength struct {
	gorm.Model
	Name string `json:"name"`
}

type PetFurColor struct {
	gorm.Model
	Name string `json:"name"`
}

type PetSize struct {
	gorm.Model
	Name string `json:"name"`
}

type PetAge struct {
	gorm.Model
	Name string `json:"name"`
}

type CreatePetDTO struct {
	Name            string `json:"name" mapstructure:"name" binding:"required,min=3,max=60"`
	FoundAddress    string `json:"foundAddress" mapstructure:"foundAddress" binding:"min=3,max=255"`
	LastSeenAddress string `json:"lastSeenAddress" mapstructure:"lastSeenAddress" binding:"min=3,max=255"`
	PetTypeID       uint   `json:"typeID" mapstructure:"typeID" binding:"required,min=1,max=60"`
	PetSpecieID     uint   `json:"specieID" mapstructure:"specieID" binding:"required,min=1,max=60"`
	PetBreedID      uint   `json:"breedID" mapstructure:"breedID" binding:"required,min=1,max=60"`
	PetFurLengthID  uint   `json:"furLengthID" mapstructure:"furLengthID" binding:"required,min=1,max=60"`
	PetFurColorID   uint   `json:"furColorID" mapstructure:"furColorID" binding:"required,min=1,max=60"`
	PetSizeID       uint   `json:"sizeID" mapstructure:"sizeID" binding:"required,min=1,max=60"`
	PetAgeID        uint   `json:"ageID" mapstructure:"ageID" binding:"required,min=1,max=60"`
}

type FindPetsDTO struct {
	shared.FindDTO
	Name            string `form:"name" mapstructure:"name" binding:"omitempty,required,min=3,max=60"`
	FoundAddress    string `form:"foundAddress" mapstructure:"foundAddress" binding:"omitempty,min=3,max=255"`
	LastSeenAddress string `form:"lastSeenAddress" mapstructure:"lastSeenAddress" binding:"omitempty,min=3,max=255"`
	PetTypeID       string `form:"typeID" mapstructure:"typeID" binding:"omitempty,required,min=1,max=60"`
	PetSpecieID     string `form:"specieID" mapstructure:"specieID" binding:"omitempty,required,min=1,max=60"`
	PetBreedID      string `form:"breedID" mapstructure:"breedID" binding:"omitempty,required,min=1,max=60"`
	PetFurLengthID  string `form:"furLengthID" mapstructure:"furLengthID" binding:"omitempty,required,min=1,max=60"`
	PetFurColorID   string `form:"furColorID" mapstructure:"furColorID" binding:"omitempty,required,min=1,max=60"`
	PetSizeID       string `form:"sizeID" mapstructure:"sizeID" binding:"omitempty,required,min=1,max=60"`
	PetAgeID        string `form:"ageID" mapstructure:"ageID" binding:"omitempty,required,min=1,max=60"`
}

func (dto *FindPetsDTO) AddFilterToStatement(query *gorm.DB) (err error) {
	if dto.Name != "" {
		err = filter.AddFilterToStatement(query, "name", dto.Name)
		if err != nil {
			return fmt.Errorf("add name filter: %w", err)
		}
	}

	if dto.FoundAddress != "" {
		err = filter.AddFilterToStatement(query, "foundAddress", dto.FoundAddress)
		if err != nil {
			return fmt.Errorf("add found address filter: %w", err)
		}
	}

	if dto.LastSeenAddress != "" {
		err = filter.AddFilterToStatement(query, "lastSeenAddress", dto.LastSeenAddress)
		if err != nil {
			return fmt.Errorf("add last seen address filter: %w", err)
		}
	}

	if dto.PetTypeID != "" {
		err = filter.AddFilterToStatement(query, "typeId", dto.PetTypeID)
		if err != nil {
			return fmt.Errorf("add pet type filter: %w", err)
		}
	}

	if dto.PetSpecieID != "" {
		err = filter.AddFilterToStatement(query, "specieId", dto.PetSpecieID)
		if err != nil {
			return fmt.Errorf("add pet specie filter: %w", err)
		}
	}

	if dto.PetBreedID != "" {
		err = filter.AddFilterToStatement(query, "breedId", dto.PetBreedID)
		if err != nil {
			return fmt.Errorf("add pet breed filter: %w", err)
		}
	}

	if dto.PetFurLengthID != "" {
		err = filter.AddFilterToStatement(query, "furLengthId", dto.PetFurLengthID)
		if err != nil {
			return fmt.Errorf("add pet fur length filter: %w", err)
		}
	}

	if dto.PetFurColorID != "" {
		err = filter.AddFilterToStatement(query, "furColorId", dto.PetFurColorID)
		if err != nil {
			return fmt.Errorf("add pet fur color filter: %w", err)
		}
	}

	if dto.PetSizeID != "" {
		err = filter.AddFilterToStatement(query, "sizeId", dto.PetSizeID)
		if err != nil {
			return fmt.Errorf("add pet size filter: %w", err)
		}
	}

	if dto.PetAgeID != "" {
		err = filter.AddFilterToStatement(query, "ageId", dto.PetAgeID)
		if err != nil {
			return fmt.Errorf("add pet age filter: %w", err)
		}
	}

	return err
}
