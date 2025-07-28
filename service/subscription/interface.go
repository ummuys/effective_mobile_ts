package service

import "github.com/ummuys/effective_mobile_ts/models"

type SubsService interface {
	CreateSubs(subsJSON *models.SubsRequest) error
	GetSubs(userID string) (*models.SubsResponse, error)
	DeleteSubs(userID string) error
	GetAllSubs() ([]models.SubsResponse, error)
	GetSumOfSubs(userID string, serviceName string, startDate string, endDate string) (int, error)
	UpdateSubs(subsJSON *models.SubsRequest) error
}
