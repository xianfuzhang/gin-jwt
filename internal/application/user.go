package application

import (
	"test/v2/internal/entities"
)

func CreateUser(userRepo entities.UserRepository, user *entities.User) error {
	return userRepo.Create(user)
}

func GetUserByName(userRepo entities.UserRepository, userName string) (entities.User, error) {
	return userRepo.GetByName(userName)
}
