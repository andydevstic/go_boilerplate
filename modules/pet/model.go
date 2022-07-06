package pet

import (
	"gorm.io/gorm"
)

type Pet struct {
	gorm.Model
	Name            string `json:"name"`
	FoundAddress    string `json:"foundAddress"`
	LastSeenAddress string `json:"lastSeenAddress"`
	PetTypeID       uint   `json:"typeID"`
	PetType         PetType
	PetSpecieID     uint `json:"specieID"`
	PetSpecie       PetSpecie
	PetBreedID      uint `json:"breedID"`
	PetBreed        PetBreed
	PetFurLengthID  uint `json:"furLengthID"`
	PetFurLength    PetFurLength
	PetFurColorID   uint `json:"furColorID"`
	PetFurColor     PetFurColor
	PetSizeID       uint `json:"sizeID"`
	PetSize         PetSize
	PetAgeID        uint `json:"ageID"`
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
