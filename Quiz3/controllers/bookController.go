package controllers

import (
	"Quiz3/library"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllBook(ctx *gin.Context){
	if ctx.Request.Method == "GET" {

		result := library.GetAllBuku()

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

func CreateBook(ctx *gin.Context){
	if ctx.Request.Method == "POST" {
		var book  library.Book
		var errorBook []library.Error

		if err := ctx.ShouldBindJSON(&book); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		if book.ReleaseYear < 1980 {
			result := library.Error{"Tahun Rilis Minimal 1980"}.Validate()
			errorBook = append(errorBook, result)
		}else if book.ReleaseYear > 2021 {
			result := library.Error{"Tahun Rilis Maximal 2021"}.Validate()
			errorBook = append(errorBook, result)
		}
		
		if _, err := url.ParseRequestURI(book.ImageUrl); err != nil {
			result := library.Error{"Gambar Harus Berupa URL"}.Validate()
			errorBook = append(errorBook, result)
		}

		if len(errorBook) > 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Status"  : false,
				"Error"   : errorBook,
			})
			return
		}

		if book.TotalPage <= 100 {
			book.Thickness = "tipis"
		}else if book.TotalPage < 200 && book.TotalPage > 100 {
			book.Thickness = "sedang"
		}else if book.TotalPage >= 201 {
			book.Thickness = "tebal"
		}else {
			book.Thickness = "tidak ada"
		}

		result 	:= library.CreateBuku(book)

		ctx.JSON(http.StatusOK, gin.H{
			"Status"	: true,
			"message"	: "Book successfully added",
			"result"	: result,
		})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusMethodNotAllowed, gin.H{
		"status":  true,
		"message": "Method Salah",
	})
}

func UpdateBook(ctx *gin.Context){
	if ctx.Request.Method == "PUT" {
		var book  library.Book
		var errorBook []library.Error

		if err := ctx.ShouldBindJSON(&book); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		if book.ReleaseYear < 1980 {
			result := library.Error{"Tahun Rilis Minimal 1980"}.Validate()
			errorBook = append(errorBook, result)
		}else if book.ReleaseYear > 2021 {
			result := library.Error{"Tahun Rilis Maximal 2021"}.Validate()
			errorBook = append(errorBook, result)
		}
		
		if _, err := url.ParseRequestURI(book.ImageUrl); err != nil {
			result := library.Error{"Gambar Harus Berupa URL"}.Validate()
			errorBook = append(errorBook, result)
		}

		if len(errorBook) > 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Status"  : false,
				"Error"   : errorBook,
			})
			return
		}

		if book.TotalPage <= 100 {
			book.Thickness = "tipis"
		}else if book.TotalPage < 200 && book.TotalPage > 100 {
			book.Thickness = "sedang"
		}else if book.TotalPage >= 201 {
			book.Thickness = "tebal"
		}else {
			book.Thickness = "tidak ada"
		}

		id 			:= ctx.Param("id")
		book.Id, _ 	= strconv.Atoi(id)

		library.UpdateBuku(book)

		ctx.JSON(http.StatusOK, gin.H{
			"Status"	: true,
			"message"	: "Book successfully updated",
		})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusMethodNotAllowed, gin.H{
		"status":  true,
		"message": "Method Salah",
	})
}

func DeleteBook(ctx *gin.Context){
	if ctx.Request.Method == "DELETE" {
		var book  library.Book

		id 			:= ctx.Param("id")
		book.Id, _	= strconv.Atoi(id)

		library.DeleteBuku(book)

		ctx.JSON(http.StatusOK, gin.H{
			"Status"	: true,
			"message"	: "Book successfully deleted",
		})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusMethodNotAllowed, gin.H{
		"status":  true,
		"message": "Method Salah",
	})
}