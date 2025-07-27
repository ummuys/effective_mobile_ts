package service

import "github.com/ummuys/effective_mobile_ts/models"

type SubsService interface {
	CreateSubs(subsJSON *models.CreateSubsRequest) error
	GetSubs(userID string) (*models.GetSubsResponse, error)
	DeleteSubs(userID string) error
	GetAllSubs() ([]models.GetSubsResponse, error)
	UpdateSubs(subsJSON *models.CreateSubsRequest) error
}
