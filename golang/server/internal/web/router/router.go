package router

import (
	"github.com/candy12t/server/internal/repository"
	"github.com/candy12t/server/internal/usecase"
	"github.com/candy12t/server/internal/web/handler"
	"github.com/gin-gonic/gin"
)

func Initialize() *gin.Engine {
	router := gin.New()
	router.Use(gin.ErrorLogger())
	router.Use(gin.Recovery())

	repoUser := repository.NewUser()
	usUser := usecase.NewUser(repoUser)
	handler.NewUser(router, usUser)

	return router
}
