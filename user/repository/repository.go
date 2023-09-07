package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/LayssonENS/go-FastHTTP-api/domain"
)

type userRepository struct {
	Conn *sql.DB
}

func NewUserRepository(conn *sql.DB) domain.UserRepository {
	return &userRepository{conn}
}

func (u *userRepository) GetByID(id int64) (domain.User, error) {
	var user domain.User

	query := `SELECT id, name FROM users WHERE id = $1`
	err := u.Conn.QueryRow(query, id).Scan(&user.ID, &user.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, fmt.Errorf("user with ID %d not found", id)
		}
		return domain.User{}, err
	}

	return user, nil
}

func (u *userRepository) CreateUser(user *domain.User) error {
	query := `
		INSERT INTO users (name)
		VALUES ($1)
		RETURNING id;
	`

	err := u.Conn.QueryRowContext(context.Background(), query, user.Name).Scan(&user.ID)
	if err != nil {
		return err
	}

	return nil
}
