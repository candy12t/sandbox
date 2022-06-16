package service

import (
	"github.com/candy12t/api-server/internal/domain/entity"
	"github.com/candy12t/api-server/internal/usecase/form"
	"github.com/candy12t/api-server/internal/usecase/repository"
	"github.com/candy12t/api-server/internal/usecase/usecase"
)

type UserUsecase struct {
	userRepository repository.User
}

var _ usecase.User = &UserUsecase{}

func NewUserUsecase(userRepository repository.User) usecase.User {
	return &UserUsecase{
		userRepository: userRepository,
	}
}

func (uu *UserUsecase) CreateUser(f *form.CreateUserInputData) (*form.UserOutputData, error) {
	user := entity.NewUser(f.Name)
	result, err := uu.userRepository.Save(user)
	if err != nil {
		return nil, err
	}
	return form.NewCreateUserOutputData(result), nil
}

func (uu *UserUsecase) GetUser(id int) (*form.UserOutputData, error) {
	user, err := uu.userRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return form.NewCreateUserOutputData(user), nil
}
