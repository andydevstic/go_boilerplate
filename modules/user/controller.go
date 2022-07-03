package user

type IUserController interface {
}

type UserController struct {
	service IUserService
}

func NewController(service IUserService) IUserController {
	return &UserController{
		service: service,
	}
}
