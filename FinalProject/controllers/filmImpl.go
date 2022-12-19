package controllers

import (
	"fmt"
	"net/http"
	"FinalProject/models"
	"FinalProject/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FilmControllerImpl struct {
	filmService services.FilmService
}

func NewFilmController(filmService services.FilmService) FilmController {
	return &FilmControllerImpl{
		filmService: filmService,
	}
}

func (f *FilmControllerImpl) AddFilm(ctx *gin.Context){
	var addFilm models.AddFilm

	if err := ctx.ShouldBindJSON(&addFilm); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	
	film, err := f.filmService.AddFilm(&addFilm)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, film)
}

func (f *FilmControllerImpl) GetFilm(ctx *gin.Context) {
	films, err := f.filmService.GetFilm()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, films)
}

func (f *FilmControllerImpl) GetFilmById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}

	film, err := f.filmService.GetFilmById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, film)
}

func (f *FilmControllerImpl) UpdateFilm(ctx *gin.Context) {
	var film models.Film

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}

	if err := ctx.ShouldBindJSON(&film); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	count, err := f.filmService.UpdateFilm(id, &film)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	message := fmt.Sprintf("Updated data amount %d", count)
	ctx.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

func (f *FilmControllerImpl) DeleteFilm(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}

	count, err := f.filmService.DeleteFilm(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	message := fmt.Sprintf("Deleted data amount %d", count)
	ctx.JSON(http.StatusOK, gin.H{
		"message": message,
	})	
}


