package controllers

import (
	"fmt"
	"net/http"
	"FinalProject/models"
	"FinalProject/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TicketControllerImpl struct {
	ticketService services.TicketService
}

func NewTicketController(ticketService services.TicketService) TicketController {
	return &TicketControllerImpl{
		ticketService: ticketService,
	}
}

func (t *TicketControllerImpl) AddTicket(ctx *gin.Context){
	var addTicket models.AddTicket

	if err := ctx.ShouldBindJSON(&addTicket); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	
	film, err := t.ticketService.AddTicket(&addTicket)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, film)
}

func (t *TicketControllerImpl) GetTicket(ctx *gin.Context) {
	films, err := t.ticketService.GetTicket()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, films)
}

func (t *TicketControllerImpl) GetTicketById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}

	film, err := t.ticketService.GetTicketById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, film)
}

func (t *TicketControllerImpl) UpdateTicket(ctx *gin.Context) {
	var ticket models.Ticket

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}

	if err := ctx.ShouldBindJSON(&ticket); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	count, err := t.ticketService.UpdateTicket(id, &ticket)
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

func (t *TicketControllerImpl) DeleteTicket(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}

	count, err := t.ticketService.DeleteTicket(id)
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


