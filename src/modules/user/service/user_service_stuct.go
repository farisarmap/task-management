package services

type UserService struct {
	CreateUserService CreateUserService
}

func NewUserService(createUserService CreateUserService) *UserService {
	return &UserService{
		CreateUserService: createUserService,
	}
}
