package form

import (
	"time"

	"github.com/candy12t/api-server/internal/domain/entity"
)

type CreateUserInputData struct {
	Name string `json:"name" binding:"required"`
}

type UserOutputData struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewCreateUserOutputData(user *entity.User) *UserOutputData {
	return &UserOutputData{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func NewUsers(users []*entity.User) []*UserOutputData {
	outputUsers := make([]*UserOutputData, 0, len(users))
	for _, user := range users {
		u := NewCreateUserOutputData(user)
		outputUsers = append(outputUsers, u)
	}
	return outputUsers
}
