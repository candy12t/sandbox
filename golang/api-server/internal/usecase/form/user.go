package form

import (
	"time"

	"github.com/candy12t/api-server/internal/domain/entity"
)

type CreateUserInputData struct {
	Name string `json:"name" binding:"required"`
}

type OutputUser struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewOutputUser(user *entity.User) *OutputUser {
	return &OutputUser{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func NewOutputUsers(users []*entity.User) []*OutputUser {
	outputUsers := make([]*OutputUser, 0, len(users))
	for _, user := range users {
		u := NewOutputUser(user)
		outputUsers = append(outputUsers, u)
	}
	return outputUsers
}
