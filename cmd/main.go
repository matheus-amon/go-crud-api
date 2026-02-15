package main

import (
	"go-crud-api/controller"
	"go-crud-api/db"
	"go-crud-api/repository"
	"go-crud-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	UserRepository := repository.NewUserRepository(dbConnection)
	UserUseCase := usecase.NewUserUseCase(UserRepository)
	UserController := controller.NewUserController(UserUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/users", UserController.GetUsers)
	server.POST("/users", UserController.CreateUser)
	server.GET("/users/:userId", UserController.GetUsersByID)

	server.Run(":8000")

}
