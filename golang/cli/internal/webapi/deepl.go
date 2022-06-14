package webapi

import (
	"github.com/candy12t/cli/internal/adapter/webapi"
	"github.com/candy12t/cli/internal/entity"
)

type DeepL struct {
}

var _ webapi.DeepL = &DeepL{}

func New() *DeepL {
	return &DeepL{}
}

func (d *DeepL) Translate(*entity.Translation) (*entity.Translation, error) {
	return &entity.Translation{}, nil
}
