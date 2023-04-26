package repository

import (
	"errors"
	"test/v2/internal/adapter/sqlite"
	"test/v2/internal/adapter/sqlite/models"
	"test/v2/internal/entities"

	"gorm.io/gorm"
)

type User struct {
}

// GetByName implements entities.UserRepository
func (*User) GetByName(name string) (entities.User, error) {
	var user = entities.User{}
	err := sqlite.DB.Where("name=?", name).First(&user).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return user, gorm.ErrRecordNotFound
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
		return err
	}
	return nil
}

// Delete implements entities.UserRepository
func (u *User) Delete(userId int64) error {
	if err := sqlite.DB.Unscoped().Delete(&models.User{}, userId).Error; err != nil {
		return err
	}
	return nil
}

// Fetch implements entities.UserRepository
func (u *User) Fetch(limit, offset int64) ([]entities.User, error) {
	var users []entities.User
	err := sqlite.DB.Limit(int(limit)).Offset(int(offset)).Find(&users).Error
	if err != nil {
		return users, err
	}
	return users, nil
}

// GetById implements entities.UserRepository
func (u *User) GetById(userId int32) (entities.User, error) {
	var user entities.User
	err := sqlite.DB.First(&user, userId).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return user, gorm.ErrRecordNotFound
	}
	return user, nil
}

// Update implements entities.UserRepository
func (u *User) Update(user *entities.User) error {
	if err := sqlite.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) UpdateUserPassword(user *entities.User) error {
	if err := sqlite.DB.Model(&entities.User{}).Where("name = ?", user.Name).Update("password", user.Password).Error; err != nil {
		return err
	}
	return nil
}

var _ entities.UserRepository = &User{}
