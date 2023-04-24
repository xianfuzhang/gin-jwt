package repository

import (
	"test/v2/internal/adapter/sqlite"
	"test/v2/internal/adapter/sqlite/models"
	"test/v2/internal/entities"
)

type User struct {
}

// ChangePassword implements entities.UserRepository
func (*User) ChangePassword(user *entities.User, newPassword string) error {
	panic("unimplemented")
}

// ResetPassword implements entities.UserRepository
func (*User) ResetPassword(user *entities.User) error {
	panic("unimplemented")
}

// GetByName implements entities.UserRepository
func (*User) GetByName(name string) (entities.User, error) {
	var user = entities.User{}
	if err := sqlite.DB.Where("name=?", name).First(&user).Error; err != nil {
		panic(err)
	}
	return user, nil
}

// Create implements entities.UserRepository
func (u *User) Create(user *entities.User) error {
	model := models.User{
		Name:     user.Name,
		Password: user.Password,
	}
	if err := sqlite.DB.Create(&model).Error; err != nil {
		panic(err)
	}
	return nil
}

// Delete implements entities.UserRepository
func (u *User) Delete(userId int64) error {
	if err := sqlite.DB.Delete(&entities.User{}, userId).Error; err != nil {
		panic(err)
	}
	return nil
}

// Fetch implements entities.UserRepository
func (u *User) Fetch(num int64) ([]entities.User, error) {
	panic("unimplemented")
}

// GetById implements entities.UserRepository
func (u *User) GetById(userId int64) (entities.User, error) {
	var user entities.User
	if err := sqlite.DB.First(&user, userId).Error; err != nil {
		panic(err)
	}
	return user, nil
}

// Update implements entities.UserRepository
func (u *User) Update(user *entities.User) error {
	if err := sqlite.DB.Save(&user).Error; err != nil {
		panic(err)
	}
	return nil
}

var _ entities.UserRepository = &User{}
