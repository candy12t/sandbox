package handler

import (
	"net/http"
	"strconv"

	"github.com/candy12t/api-server/internal/usecase/form"
	"github.com/candy12t/api-server/internal/usecase/usecase"
	"github.com/gin-gonic/gin"
)

type User struct {
	userUsecase usecase.User
}

func NewUser(userUsecase usecase.User) *User {
	return &User{
		userUsecase: userUsecase,
	}
}

func (u *User) CreateUser(c *gin.Context) {
	f := new(form.CreateUserInputData)
	if err := c.ShouldBindJSON(f); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user, err := u.userUsecase.CreateUser(f)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "create user is success!!", "result": user})
}

func (u *User) GetUser(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user, err := u.userUsecase.GetUser(id)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, user)
}
