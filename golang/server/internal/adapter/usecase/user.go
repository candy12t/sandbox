package usecase

import "github.com/candy12t/server/internal/entity"

type User interface {
	GetDetail(id int) (*entity.User, error)
	Create(name string) *entity.User
	Index() []*entity.User
}
