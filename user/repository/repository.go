package repository

import (
	"database/sql"

	"github.com/LayssonENS/go-FastHTTP-api/domain"
)

type userRepository struct {
	Conn *sql.DB
}

func NewUserRepository(conn *sql.DB) domain.UserRepository {
	return &userRepository{conn}
}

func (u *userRepository) GetByID(id int64) (domain.User, error) {

	return domain.User{
		ID:        0,
		Name:      "Test",
		CreatedAt: "test",
		UpdatedAt: "test",
	}, nil

}
