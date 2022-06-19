//go:generate mockgen -source=$GOFILE -destination=../mock_$GOPACKAGE/$GOFILE
package repository

import "github.com/candy12t/api-server/internal/domain/entity"

type Tag interface {
	FindAll() ([]*entity.Tag, error)
	FindById(id int) (*entity.Tag, error)
	Save(tag *entity.Tag) (int, error)
	Update(tag *entity.Tag) (int, error)
	Delete(tag *entity.Tag) error
}
