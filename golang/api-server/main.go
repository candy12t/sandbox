package main

import (
	"github.com/candy12t/api-server/internal/infrastructure/api/handler"
	"github.com/candy12t/api-server/internal/infrastructure/api/router"
	"github.com/candy12t/api-server/internal/infrastructure/api/server"
	"github.com/candy12t/api-server/internal/infrastructure/db"
	"github.com/candy12t/api-server/internal/usecase/service"
)

func main() {
	r := router.Initalize()

	userRepository := db.NewUserRepository()
	userUsecase := service.NewUserUsecase(userRepository)
	userHandler := handler.NewUser(userUsecase)

	userGroup := r.Group("/users")
	userGroup.POST("", userHandler.CreateUser)
	userGroup.GET(":id", userHandler.GetUser)
	userGroup.GET("", userHandler.GetUsers)
	userGroup.PUT(":id", userHandler.UpdateUser)
	userGroup.DELETE(":id", userHandler.DeleteUser)

	server.ListenAndServe(r)
}
