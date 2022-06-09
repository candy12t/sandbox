package handler

import (
	"net/http"
	"strconv"

	"github.com/candy12t/server/internal/adapter/usecase"
	"github.com/gin-gonic/gin"
)

type User struct {
	usecase usecase.User
}

func NewUser(r *gin.Engine, usecase usecase.User) {
	handler := &User{
		usecase: usecase,
	}

	r.GET("/users", handler.Index)
	r.GET("/users/", handler.Index)
	r.POST("/users", handler.Create)
	r.GET("/users/:id", handler.Show)
	r.PUT("/users/:id", handler.Update)
	r.DELETE("/users/:id", handler.Delete)
}

func (u *User) Index(c *gin.Context) {
	users := u.usecase.Index()
	c.JSON(http.StatusOK, users)
}

func (u *User) Show(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user, err := u.usecase.GetDetail(id)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (u *User) Create(c *gin.Context) {
	name := c.PostForm("name")
	user := u.usecase.Create(name)
	c.JSON(http.StatusOK, user)
}

func (u *User) Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "update"})
}

func (u *User) Delete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "delete"})
}
