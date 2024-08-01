package service

import (
	"github.com/anandaadityap/algoquill-backend/internal/app/models/entity"
	"github.com/anandaadityap/algoquill-backend/internal/app/models/web/request"
	"github.com/anandaadityap/algoquill-backend/internal/app/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(input request.RegisterUserInput) (entity.User, error)
}
type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *userService {
	return &userService{userRepository}
}

func (us *userService) RegisterUser(input request.RegisterUserInput) (entity.User, error) {
	user := entity.User{}
	user.Name = input.Name
	user.Email = input.Email
	PasswordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(PasswordHash)

	newUser, err := us.userRepository.Save(user)
	if err != nil {
		return user, err
	}

	return newUser, nil
}
