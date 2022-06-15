package repository

import "github.com/candy12t/api-server/internal/domain/entity"

type User interface {
	Save(user *entity.User) (*entity.User, error)
}
