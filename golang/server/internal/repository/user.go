package repository

import (
	"fmt"

	"github.com/candy12t/server/internal/adapter/repository"
	"github.com/candy12t/server/internal/entity"
)

type User struct{}

var _ repository.User = &User{}

var cacheUsers []*entity.User

func NewUser() *User {
	return &User{}
}

func (u *User) FindById(id int) (*entity.User, error) {
	for _, v := range cacheUsers {
		if v.ID == id {
			return v, nil
		}
	}
	return nil, fmt.Errorf("Not found user by id: %v", id)
}

func (u *User) FindAll() []*entity.User {
	return cacheUsers
}

func (u *User) Create(name string) *entity.User {
	user := entity.NewUser(name)
	cacheUsers = append(cacheUsers, user)
	return user
}
