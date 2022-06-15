package shared

type FindDTO struct {
	Limit  uint `form:"limit" binding:"min=1"`
	Offset uint `form:"offset" binding:"min=0"`
}

type ChangeLog struct {
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

type FindResponseDTO struct {
	Docs  interface{} `json:"docs"`
	Total int         `json:"total"`
}
