package routers

import (
	"Quiz3/controllers"
	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	authMiddleware := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		"admin"		: "password",
		"editor"	: "secret",
	}))

	// Route Bangun Datar
	router.GET("/bangun-datar/segitiga-sama-sisi", controllers.SegitigaSamaSisi)
	router.GET("/bangun-datar/jajar-genjang", controllers.JajarGenjang)
	router.GET("/bangun-datar/persegi", controllers.Persegi)
	router.GET("/bangun-datar/persegi-panjang", controllers.PersegiPanjang)
	router.GET("/bangun-datar/lingkaran", controllers.Lingkaran)

	// Route Categories
	router.GET("/categories", controllers.GetAllCategory)
	authMiddleware.POST("/categories", controllers.CreateCategory)

	authMiddleware.PUT("/categories/:id", controllers.UpdateCategory)
	authMiddleware.DELETE("/categories/:id", controllers.DeleteCategory)

	// Route Books
	router.GET("/books", controllers.GetAllBook)
	authMiddleware.POST("/books", controllers.CreateBook)
	authMiddleware.PUT("/books/:id", controllers.UpdateBook)
	authMiddleware.DELETE("/books/:id", controllers.DeleteBook)


	return router
}