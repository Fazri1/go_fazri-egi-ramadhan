package usecase

import (
	"belajar-go-echo/model"
	"belajar-go-echo/repository"
)

type UserUsecase interface {
	Create(payload model.User) error
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) *userUsecase {
	return &userUsecase{userRepository: userRepo}
}

func (u *userUsecase) Create(payload model.User) error {
	var user model.User
	if err := u.userRepository.Create(user); err != nil {
		return err
	}

	return nil
}
