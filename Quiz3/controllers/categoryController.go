package controllers

import (
	"net/http"
	"Quiz3/library"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetAllCategory(ctx *gin.Context){
	if ctx.Request.Method == "GET" {

		result := library.GetAllKategori()

		ctx.JSON(http.StatusOK, gin.H{
			"Status"	: true,
			"Message"	: "Successfully Get Data",
			"data"		: result,
		})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusMethodNotAllowed, gin.H{
		"status":  true,
		"message": "Method Salah",
	})
}

func CreateCategory(ctx *gin.Context){

	if ctx.Request.Method == "POST" {
		var category library.Category

		if err := ctx.ShouldBindJSON(&category); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		result 	:= library.CreateKategori(category)

		ctx.JSON(http.StatusOK, gin.H{
			"Status"	: true,
			"message"	: "Category successfully added",
			"result"	: result,
		})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusMethodNotAllowed, gin.H{
		"status":  true,
		"message": "Method Salah",
	})
}

func UpdateCategory(ctx *gin.Context){
	if ctx.Request.Method == "PUT" {
		var category library.Category

		id := ctx.Param("id")

		if err := ctx.ShouldBindJSON(&category); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		category.Id, _ = strconv.Atoi(id)

		library.UpdateKategori(category)

		ctx.JSON(http.StatusOK, gin.H{
			"Status"	: true,
			"message"	: "Category successfully udpated",
		})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusMethodNotAllowed, gin.H{
		"status":  true,
		"message": "Method Salah",
	})
}

func DeleteCategory(ctx *gin.Context){
	if ctx.Request.Method == "DELETE" {
		var category library.Category

		id := ctx.Param("id")

		category.Id, _ = strconv.Atoi(id)

		library.DeleteKategori(category)

		ctx.JSON(http.StatusOK, gin.H{
			"Status"	: true,
			"message"	: "Category successfully deleted",
		})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusMethodNotAllowed, gin.H{
		"status":  true,
		"message": "Method Salah",
	})
}