package services

import "FinalProject/models"

type StudioService interface {
	AddStudio(studio *models.AddStudio) (*models.Studio, error)
	GetStudio() (*[]models.Studio, error)
	GetStudioById(id int) (*models.Studio, error)
	UpdateStudio(id int, studio *models.Studio) (int, error)
	DeleteStudio(id int) (int, error)
}