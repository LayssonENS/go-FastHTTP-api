package usecase

import (
	"time"

	"github.com/LayssonENS/go-FastHTTP-api/domain"
)

type userUseCase struct {
	userRepo       domain.UserRepository
	contextTimeout time.Duration
}

func NewUserUseCase(a domain.UserUseCase, timeout time.Duration) domain.UserUseCase {
	return &userUseCase{
		userRepo:       a,
		contextTimeout: timeout,
	}
}

func (u *userUseCase) GetByID(id int64) (domain.User, error) {
	//TODO implement me
	panic("implement me")
}
