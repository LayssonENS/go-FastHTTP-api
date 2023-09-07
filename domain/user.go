package domain

type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type UserUseCase interface {
	GetByID(id int64) (User, error)
	CreateUser(user *User) error
}

type UserRepository interface {
	GetByID(id int64) (User, error)
	CreateUser(user *User) error
}
