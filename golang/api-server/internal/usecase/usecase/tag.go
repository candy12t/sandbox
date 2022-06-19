package usecase

import (
	"github.com/candy12t/api-server/internal/usecase/form"
)

type Tag interface {
	CreateTag(*form.CreateTagParams) (*form.OutputTag, error)
	GetTag(id int) (*form.OutputTag, error)
	GetTags() ([]*form.OutputTag, error)
	UpdateTag(*form.UpdateTagParams) (*form.OutputTag, error)
	DeleteTag(id int) error
}
