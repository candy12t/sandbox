package db

import (
	"fmt"

	"github.com/candy12t/api-server/internal/domain/entity"
	"github.com/candy12t/api-server/internal/usecase/repository"
)

type UserRepository struct {
}

var _ repository.User = &UserRepository{}

var cacheUsers []*entity.User

var id int

func NewUserRepository() repository.User {
	return &UserRepository{}
}

func (ur *UserRepository) Save(user *entity.User) (*entity.User, error) {
	id++
	user.ID = id
	cacheUsers = append(cacheUsers, user)
	return user, nil
}

func (ur *UserRepository) FindById(id int) (*entity.User, error) {
	for _, user := range cacheUsers {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, fmt.Errorf("Not fount user by: %v", id)
}
