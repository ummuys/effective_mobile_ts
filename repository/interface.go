package repository

import "github.com/ummuys/effective_mobile_ts/models"

type Database interface {
	CreateSubs(subsInfo models.Subs) error
	GetSubs(userID string) (*models.SubsDB, error)
	DeleteSubs(userID string) error
	GetAllSubs() ([]models.SubsDB, error)
	UpdateSubs(subsInfo models.Subs) error
	CheckUserExists(userID string) (bool, error)
}
