package shared

type FindDTO struct {
	Limit  uint `form:"limit" binding:"min=1"`
	Offset uint `form:"offset" binding:"min=0"`
}

type FindResponseDTO struct {
	Docs  interface{} `json:"docs"`
	Total int         `json:"total"`
}

type UserAuthPayload struct {
	Id       uint   `json:"id" mapstructure:"id"`
	Email    string `json:"email" mapstructure:"email"`
	Username string `json:"username" mapstructure:"username"`
	Name     string `json:"name" binding:"max=100"`
	Type     uint8  `json:"type" mapstructure:"type"`
	Status   uint8  `json:"status" mapstructure:"status"`
}
