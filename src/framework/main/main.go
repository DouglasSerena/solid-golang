package main

import (
	"github.com/DouglasSerena/solid-go/src/app/presentation/controllers"
	"github.com/DouglasSerena/solid-go/src/app/repositories"
	"github.com/DouglasSerena/solid-go/src/app/usecase"
	"github.com/DouglasSerena/solid-go/src/framework/main/routes"
	"github.com/DouglasSerena/solid-go/src/framework/main/server"
	"github.com/DouglasSerena/solid-go/src/framework/utils"
)

func main() {
	server := server.NewServer()
	db := utils.ConnectDB()

	userRepository := repositories.UserRepositoryDb{Db: db}
	userUsecase := usecase.UserUsecase{UserRepository: &userRepository}
	userController := controllers.UserController{UserUsecase: &userUsecase}
	routes.UserSetupRoutes(&userController, server.Route)

	server.Route.Run("localhost:3333")
}
