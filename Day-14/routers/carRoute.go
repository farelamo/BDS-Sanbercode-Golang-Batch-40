package routers

import (
	"Day-14/controllers"
	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/cars", controllers.CreateCar)
	return router
}