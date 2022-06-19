//go:generate mockgen -source=$GOFILE -destination=../mock_$GOPACKAGE/$GOFILE
package repository

import "github.com/candy12t/api-server/internal/domain/entity"

type User interface {
	FindAll() ([]*entity.User, error)
	FindById(id int) (*entity.User, error)
	Save(user *entity.User) (int, error)
	Update(user *entity.User) (int, error)
	Delete(user *entity.User) error
}
