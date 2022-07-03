package authentication

import "github.com/andydevstic/boilerplate-backend/shared"

type RegisterUserDTO struct {
	Email       string `json:"email" binding:"required,email,min=3,max=100"`
	Username    string `json:"username" binding:"required,min=3,max=100"`
	Password    string `json:"password" binding:"required,min=6,max=60"`
	Address     string `json:"address" binding:"min=3,max=255"`
	PhoneNumber string `json:"phoneNumber" binding:"min=3,max=20"`
}

type ChangeUserPasswordDTO struct {
	Email       string `json:"email" binding:"required,email,min=3,max=100"`
	OldPassword string `json:"old_password" binding:"required,min=6,max=255"`
	NewPassword string `json:"new_password" binding:"required,min=6,max=255"`
}

type ResetUserPassword struct {
	Email string `json:"email" binding:"required,email,min=3,max=100"`
}

type LoginDTO struct {
	Email    string `json:"email" binding:"required,email,min=3,max=100"`
	Password string `json:"password" binding:"required,min=6,max=60"`
}

type LoginResponse struct {
	Token string                 `json:"token"`
	User  shared.UserAuthPayload `json:"user"`
}
