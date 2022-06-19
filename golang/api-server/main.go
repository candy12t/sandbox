package main

import (
	"log"

	"github.com/candy12t/api-server/internal/infrastructure/api/handler"
	"github.com/candy12t/api-server/internal/infrastructure/api/router"
	"github.com/candy12t/api-server/internal/infrastructure/api/server"
	"github.com/candy12t/api-server/internal/infrastructure/database"
	"github.com/candy12t/api-server/internal/usecase/service"
)

func main() {
	db, err := database.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := router.Initalize()

	userRepository := database.NewUserRepository(db)
	tagRepository := database.NewTagRepository(db)

	userUsecase := service.NewUserUsecase(userRepository)
	tagUsecase := service.NewTagUsecase(tagRepository)

	userHandler := handler.NewUser(userUsecase)
	tagHandler := handler.NewTag(tagUsecase)

	userGroup := r.Group("/users")
	userGroup.POST("", userHandler.CreateUser)
	userGroup.GET(":id", userHandler.GetUser)
	userGroup.GET("", userHandler.GetUsers)
	userGroup.PUT(":id", userHandler.UpdateUser)
	userGroup.DELETE(":id", userHandler.DeleteUser)

	tagGroup := r.Group("/tags")
	tagGroup.POST("", tagHandler.CreateTag)
	tagGroup.GET(":id", tagHandler.GetTag)
	tagGroup.GET("", tagHandler.GetTags)
	tagGroup.PUT(":id", tagHandler.UpdateTag)
	tagGroup.DELETE(":id", tagHandler.DeleteTag)

	server.ListenAndServe(r)
}
