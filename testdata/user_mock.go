package testdata

import (
	"test/v2/internal/entities"
	"test/v2/internal/types"
)

type MockUserRepo struct {
	entities.UserRepository
}

func (m *MockUserRepo) ResetPassword(user *entities.User) error {
	user.Password = types.UserResetPwd
	return nil
}

func (m *MockUserRepo) ChangePassword(user *entities.User, newPwd string) error {
	user.Password = newPwd
	return nil
}
