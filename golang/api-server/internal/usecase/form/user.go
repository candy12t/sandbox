package form

import (
	"time"

	"github.com/candy12t/api-server/internal/domain/entity"
)

type CreateUserInputData struct {
	Name string `json:"name" binding:"required"`
}

type CreateUserOutputData struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewCreateUserOutputData(user *entity.User) *CreateUserOutputData {
	return &CreateUserOutputData{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
