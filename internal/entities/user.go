package entities

type User struct {
	ID       int64
	Name     string
	Password string
}

type UserRepository interface {
	Create(user *User) error
	Update(user *User) error
	Delete(userId int64) error
	GetById(userId int64) User
	Fetch(num int64) ([]User, error)
}
