package constants

const ParsedDtoKey string = "dto"
const DbDialect string = "postgres"
const UserAuthPayload string = "user"

const InternalServerErrorMsg string = "internal server error"

const (
	StatusInActive = iota + 1
	StatusActive
	StatusDeleted
)
