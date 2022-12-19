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

	userService 		:= services.NewUserService(DB)
	userController 		:= controllers.NewUserController(userService)

	filmService 		:= services.NewFilmService(DB)
	filmController 		:= controllers.NewFilmController(filmService)

	studioService 		:= services.NewStudioService(DB)
	studioController 	:= controllers.NewStudioController(studioService)
	
	ticketService 		:= services.NewTicketService(DB)
	ticketController 	:= controllers.NewTicketController(ticketService)

	/* User Route */
	router.POST("/user", userController.AddUser)
	router.GET("/user", middleware.Authenticate, userController.GetUser)
	router.GET("/user/:id", middleware.Authenticate, userController.GetUserById)
	router.PUT("/user/:id", middleware.Authenticate, userController.UpdateUser)
	router.DELETE("/user/:id", middleware.Authenticate, userController.DeleteUser)

	/* Film Route */
	router.POST("/film", middleware.Authenticate, filmController.AddFilm)
	router.GET("/film", middleware.Authenticate, filmController.GetFilm)
	router.GET("/film/:id", middleware.Authenticate, filmController.GetFilmById)
	router.PUT("/film/:id", middleware.Authenticate, filmController.UpdateFilm)
	router.DELETE("/film/:id", middleware.Authenticate, filmController.DeleteFilm)

	/* Studio Route */
	router.POST("/studio", middleware.Authenticate, studioController.AddStudio)
	router.GET("/studio", middleware.Authenticate, studioController.GetStudio)
	router.GET("/studio/:id", middleware.Authenticate, studioController.GetStudioById)
	router.PUT("/studio/:id", middleware.Authenticate, studioController.UpdateStudio)
	router.DELETE("/studio/:id", middleware.Authenticate, studioController.DeleteStudio)

	/* Ticket Route */
	router.POST("/ticket", middleware.Authenticate, ticketController.AddTicket)
	router.GET("/ticket", middleware.Authenticate, ticketController.GetTicket)
	router.GET("/ticket/:id", middleware.Authenticate, ticketController.GetTicketById)
	router.PUT("/ticket/:id", middleware.Authenticate, ticketController.UpdateTicket)
	router.DELETE("/ticket/:id", middleware.Authenticate, ticketController.DeleteTicket)

	runWithPort := fmt.Sprintf("0.0.0.0:%s", "8000")
	router.Run(runWithPort)
}