package domain

type User struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UserUseCase interface {
	GetByID(id int64) (User, error)
}

type UserRepository interface {
	GetByID(id int64) (User, error)
}
