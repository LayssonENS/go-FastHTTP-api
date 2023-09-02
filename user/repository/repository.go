package repository

import (
	"database/sql"
	"log"

	"github.com/LayssonENS/go-FastHTTP-api/domain"
)

type userRepository struct {
	Conn *sql.DB
}

func NewUserRepository(conn *sql.DB) domain.UserRepository {
	return &userRepository{conn}
}

func (u *userRepository) GetByID(id int64) (domain.User, error) {
	db, err := sql.Open("postgres", "user=DB_USER password=DB_PASSWORD dbname=DB_NAME sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	row := db.QueryRow("SELECT id, name FROM users WHERE id = $1", id)

	user := new(domain.User)
	err = row.Scan(&user.ID, &user.Name)
	if err != nil {
		return domain.User{}, err
	}

	return *user, nil
}
