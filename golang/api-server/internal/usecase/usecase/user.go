package usecase

import (
	"github.com/candy12t/api-server/internal/usecase/form"
)

type User interface {
	CreateUser(*form.CreateUserParams) (*form.OutputUser, error)
	GetUser(id int) (*form.OutputUser, error)
	GetUsers() ([]*form.OutputUser, error)
	UpdateUser(*form.UpdateUserParams) (*form.OutputUser, error)
}
