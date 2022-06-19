package form

import (
	"time"

	"github.com/candy12t/api-server/internal/domain/entity"
)

type CreateTagParams struct {
	Name string `json:"name" binding:"required"`
}

type UpdateTagParams struct {
	ID   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type OutputTag struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewOutputTag(tag *entity.Tag) *OutputTag {
	return &OutputTag{
		ID:        tag.ID,
		Name:      tag.Name,
		CreatedAt: tag.CreatedAt,
		UpdatedAt: tag.UpdatedAt,
	}
}

func NewOutputTags(tags []*entity.Tag) []*OutputTag {
	outputTags := make([]*OutputTag, 0, len(tags))
	for _, tag := range tags {
		t := NewOutputTag(tag)
		outputTags = append(outputTags, t)
	}
	return outputTags
}
