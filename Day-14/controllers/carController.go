package controllers

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

var CarDatas = []Car{}

type Car struct {
	CarID string `json:"car_id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Price int	 `json:"price"`
}

func CreateCar(ctx *gin.Context){
	var newCar Car

	if err := ctx.ShouldBindJSON(&newCar); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return 
	}

	newCar.CarID = fmt.Sprintf("c%d", len(CarDatas)+1)
	CarDatas     = append(CarDatas, newCar)

	ctx.JSON(http.StatusCreated, gin.H{
		"car": newCar,
	})
}

func UpdateCar(ctx *gin.Context){
	carID 		:= ctx.param("carID")
	condition 	:= false 
	var updateCar Car
	if err = ctx.ShouldBindJSON(&updateCar); err != nil {
		
	}
}

