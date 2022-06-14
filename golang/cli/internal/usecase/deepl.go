package usecase

import (
	"github.com/candy12t/cli/internal/adapter/usecase"
	"github.com/candy12t/cli/internal/adapter/webapi"
	"github.com/candy12t/cli/internal/entity"
)

type DeepL struct {
	api webapi.DeepL
}

var _ usecase.DeepL = &DeepL{}

func New(api webapi.DeepL) *DeepL {
	return &DeepL{
		api: api,
	}
}

func (d *DeepL) Translate(original, sourceLanguage, targetLanguage string) (string, error) {
	t := &entity.Translation{
		Original:       original,
		SourceLanguage: sourceLanguage,
		TargetLanguage: targetLanguage,
	}
	tt, err := d.api.Translate(t)
	if err != nil {
		return "", err
	}
	return tt.Translation, nil
}
