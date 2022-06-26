package service

import (
	"fmt"
	"testing"
	"time"

	"github.com/candy12t/api-server/internal/domain/entity"
	"github.com/candy12t/api-server/internal/usecase/form"
	"github.com/candy12t/api-server/internal/usecase/mock_repository"
	"github.com/golang/mock/gomock"
)

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserRepository := mock_repository.NewMockUser(ctrl)
	mockUserRepository.EXPECT().Save(&entity.User{
		Name: "name",
	}).Return(1, nil)
	timeNow := time.Now()
	mockUserRepository.EXPECT().FindById(1).Return(&entity.User{
		ID:         1,
		Name:       "name",
		CreatedAt:  timeNow,
		UpdatedAt:  timeNow,
		DeleteMark: false,
	}, nil)

	u := NewUserUsecase(mockUserRepository)
	result, err := u.CreateUser(&form.CreateUserParams{Name: "name"})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}
