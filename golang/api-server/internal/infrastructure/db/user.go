package db

import (
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
