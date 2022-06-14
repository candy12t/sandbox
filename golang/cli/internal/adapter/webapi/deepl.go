package webapi

import "github.com/candy12t/cli/internal/entity"

type DeepL interface {
	Translate(*entity.Translation) (*entity.Translation, error)
}
