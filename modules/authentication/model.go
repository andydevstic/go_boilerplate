package authentication

type RegisterUserDTO struct {
	Email    string `json:"email" binding:"required,email,min=3,max=100"`
	Name     string `json:"name" binding:"required,min=3,max=100"`
	Password string `json:"password" binding:"required,min=6,max=60"`
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

type UserAuthPayload struct {
	Id     int    `json:"id" mapstructure:"id"`
	Email  string `json:"email" mapstructure:"email"`
	Name   string `json:"name" mapstructure:"name"`
	Type   uint8  `json:"type" mapstructure:"type"`
	Status uint8  `json:"status" mapstructure:"status"`
}
