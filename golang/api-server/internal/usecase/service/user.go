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

func (uu *UserUsecase) CreateUser(f *form.CreateUserParams) (*form.OutputUser, error) {
	user := &entity.User{
		Name: f.Name,
	}
	id, err := uu.userRepository.Save(user)
	if err != nil {
		return nil, err
	}

	result, err := uu.userRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return form.NewOutputUser(result), nil
}

func (uu *UserUsecase) GetUser(id int) (*form.OutputUser, error) {
	user, err := uu.userRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return form.NewOutputUser(user), nil
}

func (uu *UserUsecase) GetUsers() ([]*form.OutputUser, error) {
	users, err := uu.userRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return form.NewOutputUsers(users), nil
}

func (uu *UserUsecase) UpdateUser(f *form.UpdateUserParams) (*form.OutputUser, error) {
	user := &entity.User{
		ID:   f.ID,
		Name: f.Name,
	}
	id, err := uu.userRepository.Update(user)
	if err != nil {
		return nil, err
	}

	result, err := uu.userRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return form.NewOutputUser(result), nil
}

func (uu *UserUsecase) DeleteUser(id int) error {
	user := &entity.User{
		ID:         id,
		DeleteMark: true,
	}
	return uu.userRepository.Delete(user)
}
