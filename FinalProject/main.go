package main

import (
	"fmt"
	"FinalProject/config"
	"FinalProject/controllers"
	"FinalProject/middleware"
	"FinalProject/services"

	"github.com/gin-gonic/gin"
)

func main(){
	router 	:= gin.Default()
	DB 		:= config.Connect()

	userService 	:= services.NewUserService(DB)
	userController 	:= controllers.NewUserController(userService)
	router.POST("/user", userController.AddUser)
	router.GET("/user", middleware.Authenticate, userController.GetUser)
	router.GET("/user/:id", middleware.Authenticate, userController.GetUserById)
	router.PUT("/user/:id", middleware.Authenticate, userController.UpdateUser)
	router.DELETE("/user/:id", middleware.Authenticate, userController.DeleteUser)

	runWithPort := fmt.Sprintf("0.0.0.0:%s", "8000")
	router.Run(runWithPort)
}