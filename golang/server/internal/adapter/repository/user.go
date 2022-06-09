package repository

import "github.com/candy12t/server/internal/entity"

type User interface {
	FindById(id int) (*entity.User, error)
	FindAll() []*entity.User
	Create(name string) *entity.User
	// Update(id int) *entity.User
}
