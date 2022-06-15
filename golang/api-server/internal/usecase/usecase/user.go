package usecase

import (
	"github.com/candy12t/api-server/internal/usecase/form"
)

type User interface {
	CreateUser(*form.CreateUserInputData) (*form.CreateUserOutputData, error)
}
