package services

import "FinalProject/models"

type TicketService interface {
	AddTicket(ticket *models.AddTicket) (*models.Ticket, error)
	GetTicket() (*[]models.Ticket, error)
	GetTicketById(id int) (*models.Ticket, error)
	UpdateTicket(id int, ticket *models.Ticket) (int, error)
	DeleteTicket(id int) (int, error)
}