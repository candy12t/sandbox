package service

import (
	"github.com/candy12t/api-server/internal/domain/entity"
	"github.com/candy12t/api-server/internal/usecase/form"
	"github.com/candy12t/api-server/internal/usecase/repository"
	"github.com/candy12t/api-server/internal/usecase/usecase"
)

type TagUsecase struct {
	tagRepository repository.Tag
}

var _ usecase.Tag = &TagUsecase{}

func NewTagUsecase(tagRepository repository.Tag) usecase.Tag {
	return &TagUsecase{
		tagRepository: tagRepository,
	}
}

func (tu *TagUsecase) CreateTag(f *form.CreateTagParams) (*form.OutputTag, error) {
	tag := &entity.Tag{
		Name: f.Name,
	}
	id, err := tu.tagRepository.Save(tag)
	if err != nil {
		return nil, err
	}

	result, err := tu.tagRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return form.NewOutputTag(result), nil
}

func (tu *TagUsecase) GetTag(id int) (*form.OutputTag, error) {
	tag, err := tu.tagRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return form.NewOutputTag(tag), nil
}

func (tu *TagUsecase) GetTags() ([]*form.OutputTag, error) {
	tags, err := tu.tagRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return form.NewOutputTags(tags), nil
}

func (tu *TagUsecase) UpdateTag(f *form.UpdateTagParams) (*form.OutputTag, error) {
	tag := &entity.Tag{
		ID:   f.ID,
		Name: f.Name,
	}
	id, err := tu.tagRepository.Update(tag)
	if err != nil {
		return nil, err
	}

	result, err := tu.tagRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return form.NewOutputTag(result), nil
}

func (tu *TagUsecase) DeleteTag(id int) error {
	tag := &entity.Tag{
		ID:         id,
		DeleteMark: true,
	}
	return tu.tagRepository.Delete(tag)
}
