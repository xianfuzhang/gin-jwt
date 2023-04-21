package application

import (
	"test/v2/internal/entities"
)

func CreateUser(userRepo entities.UserRepository, user *entities.User) error {
	return userRepo.Create(user)
}
