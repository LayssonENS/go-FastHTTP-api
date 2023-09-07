package usecase

import (
	"github.com/LayssonENS/go-FastHTTP-api/domain"
)

type userUseCase struct {
	userRepo domain.UserRepository
}

func NewUserUseCase(a domain.UserUseCase) domain.UserUseCase {
	return &userUseCase{
		userRepo: a,
	}
}

func (u *userUseCase) GetByID(id int64) (domain.User, error) {
	user, err := u.userRepo.GetByID(id)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (u *userUseCase) CreateUser(user *domain.User) error {
	err := u.userRepo.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}
