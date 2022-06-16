package repository

import "github.com/candy12t/api-server/internal/domain/entity"

type User interface {
	Save(user *entity.User) (*entity.User, error)
	FindById(id int) (*entity.User, error)
	FindAll() ([]*entity.User, error)
	Update(user *entity.User) (*entity.User, error)
	Delete(user *entity.User) error
}
