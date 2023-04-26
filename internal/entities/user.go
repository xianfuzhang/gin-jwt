package entities

type User struct {
	ID       uint32
	Name     string
	Password string
}

type UserRepository interface {
	Create(user *User) error
	Update(user *User) error
	Delete(userId int64) error
	GetById(userId int32) (User, error)
	GetByName(name string) (User, error)
	Fetch(limit, offset int64) ([]User, error)

	UpdateUserPassword(user *User) error
}
