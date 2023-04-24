package entities

type User struct {
	Name     string
	Password string
}

type UserRepository interface {
	Create(user *User) error
	Update(user *User) error
	Delete(userId int64) error
	GetById(userId int64) (User, error)
	GetByName(name string) (User, error)
	Fetch(num int64) ([]User, error)

	UpdateUserPassword(user *User) error
}
