package routers

import (
	"Quiz3/controllers"
	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	// Route Bangun Datar
	router.GET("/bangun-datar/segitiga-sama-sisi", controllers.SegitigaSamaSisi)
	router.GET("/bangun-datar/jajar-genjang", controllers.JajarGenjang)
	router.GET("/bangun-datar/persegi", controllers.Persegi)
	router.GET("/bangun-datar/persegi-panjang", controllers.PersegiPanjang)
	router.GET("/bangun-datar/lingkaran", controllers.Lingkaran)

	// Route Categories
	router.GET("/categories", controllers.GetAllCategory)
	router.POST("/categories", controllers.CreateCategory)
	router.PUT("/categories/:categoryId", controllers.UpdateCategory)
	router.DELETE("/categories/:categoryId", controllers.DeleteCategory)

	// Route Books
	router.POST("/books", controllers.CreateBook)


	return router
}