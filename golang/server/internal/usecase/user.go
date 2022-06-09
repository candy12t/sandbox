package usecase

import (
	"github.com/candy12t/server/internal/adapter/repository"
	"github.com/candy12t/server/internal/adapter/usecase"
	"github.com/candy12t/server/internal/entity"
)

var _ usecase.User = &User{}

type User struct {
	repo repository.User
}

func NewUser(repo repository.User) *User {
	return &User{
		repo: repo,
	}
}

func (u *User) GetDetail(id int) (*entity.User, error) {
	return u.repo.FindById(id)
}

func (u *User) Index() []*entity.User {
	return u.repo.FindAll()
}

func (u *User) Create(name string) *entity.User {
	return u.repo.Create(name)
}
