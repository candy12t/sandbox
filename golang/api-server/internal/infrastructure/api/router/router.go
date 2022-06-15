package router

import (
	"github.com/gin-gonic/gin"
)

func Initalize() *gin.Engine {
	router := gin.New()
	router.Use(gin.ErrorLogger())
	return router
}
