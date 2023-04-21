package repository

import (
	"test/v2/internal/adapter/sqlite"
	"test/v2/internal/adapter/sqlite/models"
	"test/v2/internal/entities"
	"test/v2/internal/utils"
)

type User struct {
}

// Create implements entities.UserRepository
func (u *User) Create(user *entities.User) error {
	model := models.User{
		Name:     user.Name,
		Password: user.Password,
	}
	result := sqlite.DB.Create(&model)
	return utils.HandleError(result)
}

// Delete implements entities.UserRepository
func (u *User) Delete(userId int64) error {
	result := sqlite.DB.Delete(&entities.User{}, userId)
	return utils.HandleError(result)
}

// Fetch implements entities.UserRepository
func (u *User) Fetch(num int64) ([]entities.User, error) {
	panic("unimplemented")
}

// GetById implements entities.UserRepository
func (u *User) GetById(userId int64) entities.User {
	var user entities.User
	result := sqlite.DB.First(&user, userId)
	if result.Error != nil {
		panic(result.Error)
	}
	return user
}

// Update implements entities.UserRepository
func (u *User) Update(user *entities.User) error {
	result := sqlite.DB.Save(&user)
	return utils.HandleError(result)
}

var _ entities.UserRepository = &User{}
