package user

import (
	"github.com/andydevstic/boilerplate-backend/shared"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique"`
	Name     string `json:"name" gorm:"index"`
	Type     uint8  `json:"type"`
	Status   uint8  `json:"status"`
	Password string `json:"password" gorm:"not null"`
}

type FindUsersAdminDTO struct {
	shared.FindDTO
	Email  string `json:"email" binding:"email,min=3,max=60"`
	Name   string `json:"name" binding:"min=3,max=60"`
	Type   string `json:"type" binding:"min=1,max=10"`
	Status string `json:"status" binding:"min=1,max=10"`
}

type CreateUserDTO struct {
	Email    string `json:"email" binding:"required,email,min=3,max=100"`
	Name     string `json:"name" binding:"required,min=3,max=100"`
	Password string `json:"password" binding:"required,min=6,max=60"`
}

type UpsertUserAdminDTO struct {
	Email    string `json:"email" binding:"required,email,min=3,max=100"`
	Name     string `json:"name" binding:"required,min=3,max=100"`
	Password string `json:"password" binding:"min=6,max=60"`
	Type     uint8  `json:"type" binding:"required,min=1,max=10"`
	Status   uint8  `json:"status" binding:"required,min=1,max=10"`
}

type UpdateUserDTO struct {
	Name string `json:"name" binding:"required,min=3,max=100"`
}

type FindUserByEmail struct {
	Email string `json:"email" binding:"required,email,min=3,max=100"`
}
