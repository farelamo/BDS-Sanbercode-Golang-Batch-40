package controllers

import (
	"Quiz3/library"
	"github.com/gin-gonic/gin"
	"net/http"
	_"strconv"
)

func CreateBook(ctx *gin.Context){
	if ctx.Request.Method == "POST" {
		var book library.Book

		if err := ctx.ShouldBindJSON(&book); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		result 	:= library.CreateBuku(book)

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