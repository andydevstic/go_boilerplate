package constants

const ParsedDtoKey string = "dto"
const DbDialect string = "postgres"
const UserAuthPayload string = "user"

const InternalServerErrorMsg string = "Internal Server Error"

const (
	StatusInActive = iota + 1
	StatusActive
	StatusDeleted
)
