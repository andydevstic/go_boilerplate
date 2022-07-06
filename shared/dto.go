package shared

type FilterCriteria struct {
	Code     string `json:"code" binding:"required,min=2,max=30"`
	Operator string `json:"operator" binding:"required,min=2,max=30"`
	Value    any    `json:"value"`
}

type FindDTO struct {
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
