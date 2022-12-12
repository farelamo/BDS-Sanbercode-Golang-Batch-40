package controllers

import (
	"Quiz3/library"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var BangunDatar library.HitungBangunDatar

func SegitigaSamaSisi(ctx *gin.Context) {
	var result int

	if ctx.Request.Method == "GET" {

		hitung 		:= ctx.Query("hitung")
		alas, _ 	:= strconv.Atoi(ctx.Query("alas"))
		tinggi, _ 	:= strconv.Atoi(ctx.Query("tinggi"))

		BangunDatar = library.SegitigaSamaSisi{alas, tinggi}

		if hitung == "luas" {
			result = BangunDatar.Luas()
		} else {
			result = BangunDatar.Keliling()
		}

		ctx.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Successfully Calculated Data",
			"result":  result,
		})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusMethodNotAllowed, gin.H{
		"status":  true,
		"message": "Method Salah",
	})
}

func PersegiPanjang(ctx *gin.Context) {
	var result int

	if ctx.Request.Method == "GET" {

		hitung 		:= ctx.Query("hitung")
		panjang, _ 	:= strconv.Atoi(ctx.Query("panjang"))
		lebar, _ 	:= strconv.Atoi(ctx.Query("lebar"))

		BangunDatar = library.PersegiPanjang{panjang, lebar}

		if hitung == "luas" {
			result = BangunDatar.Luas()
		} else {
			result = BangunDatar.Keliling()
		}

		ctx.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Successfully Calculated Data",
			"result":  result,
		})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusMethodNotAllowed, gin.H{
		"status":  true,
		"message": "Method Salah",
	})
}

func Persegi(ctx *gin.Context) {
	var result int

	if ctx.Request.Method == "GET" {

		hitung 		:= ctx.Query("hitung")
		sisi, _ 	:= strconv.Atoi(ctx.Query("sisi"))

		BangunDatar = library.Persegi{sisi}

		if hitung == "luas" {
			result = BangunDatar.Luas()
		} else {
			result = BangunDatar.Keliling()
		}

		ctx.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Successfully Calculated Data",
			"result":  result,
		})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusMethodNotAllowed, gin.H{
		"status":  true,
		"message": "Method Salah",
	})
}

func JajarGenjang(ctx *gin.Context) {
	var result int

	if ctx.Request.Method == "GET" {

		hitung 		:= ctx.Query("hitung")
		sisi, _ 	:= strconv.Atoi(ctx.Query("sisi"))
		alas, _ 	:= strconv.Atoi(ctx.Query("alas"))
		tinggi, _ 	:= strconv.Atoi(ctx.Query("tinggi"))

		BangunDatar = library.JajarGenjang{float64 (sisi), float64 (alas), float64 (tinggi)}

		if hitung == "luas" {
			result = BangunDatar.Luas()
		} else {
			result = BangunDatar.Keliling()
		}

		ctx.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Successfully Calculated Data",
			"result":  result,
		})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusMethodNotAllowed, gin.H{
		"status":  true,
		"message": "Method Salah",
	})
}

func Lingkaran(ctx *gin.Context) {
	var result int

	if ctx.Request.Method == "GET" {

		hitung 		:= ctx.Query("hitung")
		jari, _ 	:= strconv.Atoi(ctx.Query("jari"))

		BangunDatar = library.Lingkaran{float64 (jari)}

		if hitung == "luas" {
			result = BangunDatar.Luas()
		} else {
			result = BangunDatar.Keliling()
		}

		ctx.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Successfully Calculated Data",
			"result":  result,
		})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusMethodNotAllowed, gin.H{
		"status":  true,
		"message": "Method Salah",
	})
}

