package adapter

import (
	"test/v2/internal/entities"
	"test/v2/internal/service"
	"test/v2/internal/types"
	"test/v2/testdata"
	"testing"
)

func TestResetPassword(t *testing.T) {
	mockUser := entities.User{Name: "aaa", Password: "aaaaa"}
	mockUerRepo := testdata.MockUserRepo{}

	err := service.UpdateUserPassword(&mockUerRepo, &mockUser)
	if err != nil {
		t.Fatal(err)
	}
	if mockUser.Password != types.UserResetPwd {
		t.Errorf("got new password %q, want password %q", mockUser.Password, types.UserResetPwd)
	}
}
