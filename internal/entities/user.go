package entities

import "time"

type User struct {
	Name     string
	Password string
}

type UserResponse struct {
	ID       uint
	Name     string
	Password string
	CreateAt time.Time
}

type UserRepository interface {
	Create(user *User) error
	// Update(user *User) error
	Delete(userId int64) error
	GetById(userId int32) (UserResponse, error)
	GetByName(name string) (UserResponse, error)
	Fetch(limit, offset int64) ([]UserResponse, error)

	UpdateUserPassword(user *User) error
}
