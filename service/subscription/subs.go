package service

import (
	"fmt"
	"time"

	"github.com/rs/zerolog"
	"github.com/ummuys/effective_mobile_ts/models"
	"github.com/ummuys/effective_mobile_ts/repository"
)

type subsServ struct {
	db     repository.Database
	logger *zerolog.Logger
}

func NewSubsService(db repository.Database, logger *zerolog.Logger) SubsService {
	return &subsServ{
		db:     db,
		logger: logger,
	}
}

func (fs *subsServ) CreateSubs(subsJSON *models.SubsRequest) error {
	if err := validUserID(subsJSON.UserID); err != nil {
		return err
	}

	exists, err := fs.db.CheckUserExists(subsJSON.UserID)
	if err != nil {
		fs.logger.Error().
			Str("error", err.Error()).
			Msg("problem with database")
		return err
	}

	if err := validServiceName(subsJSON.ServiceName); err != nil {
		return err
	}

	err = validPrice(subsJSON.Price)
	if err != nil {
		return err
	}

	if exists {
		return repository.ErrUserExists
	}

	startDate, err := ymtoymd(subsJSON.StartDate)
	if err != nil {
		return fmt.Errorf("bad start_date")
	}

	var endDate string
	if subsJSON.EndDate == "" {
		endDate = time.Date(9999, 12, 1, 0, 0, 0, 0, time.UTC).Format("2006-01-02")
	} else {
		endDate, err = ymtoymd(subsJSON.EndDate)
		if err != nil {
			return fmt.Errorf("bad end_date")
		}
	}

	subs := models.Subs{
		ServiceName: subsJSON.ServiceName,
		Price:       subsJSON.Price,
		UserID:      subsJSON.UserID,
		StartDate:   startDate,
		EndDate:     endDate,
	}

	if err := fs.db.CreateSubs(subs); err != nil {
		fs.logger.Error().
			Str("error", err.Error()).
			Msg("problem with database")
		return err
	}

	return nil
}

func (fs *subsServ) GetSubs(userID string) (*models.SubsResponse, error) {
	if err := validUserID(userID); err != nil {
		return nil, err
	}

	exists, err := fs.db.CheckUserExists(userID)
	if err != nil {
		fs.logger.Error().
			Str("error", err.Error()).
			Msg("problem with database")
		return nil, err
	}

	if !exists {
		return nil, repository.ErrUserDoesntExists
	}

	subsResp, err := fs.db.GetSubs(userID)
	if err != nil {
		fs.logger.Error().
			Str("error", err.Error()).
			Msg("problem with database")
		return nil, err
	}

	startDateStr := ymdtoym(subsResp.StartDate)
	endDateStr := ymdtoym(subsResp.EndDate)

	return &models.SubsResponse{
		ServiceName: subsResp.ServiceName,
		Price:       subsResp.Price,
		UserID:      subsResp.UserID,
		StartDate:   startDateStr,
		EndDate:     endDateStr,
	}, nil

}

func (fs *subsServ) DeleteSubs(userID string) error {
	if err := validUserID(userID); err != nil {
		return err
	}

	exists, err := fs.db.CheckUserExists(userID)
	if err != nil {
		fs.logger.Error().
			Str("error", err.Error()).
			Msg("problem with database")
		return err
	}

	if !exists {
		return repository.ErrUserDoesntExists
	}

	if err := fs.db.DeleteSubs(userID); err != nil {
		fs.logger.Error().
			Str("error", err.Error()).
			Msg("problem with database")
		return err
	}
	return nil
}

func (fs *subsServ) GetAllSubs() ([]models.SubsResponse, error) {
	allSubs, err := fs.db.GetAllSubs()
	if err != nil {
		fs.logger.Error().
			Str("error", err.Error()).
			Msg("problem with database")
		return nil, err
	}

	var allSubsResp []models.SubsResponse
	for _, elem := range allSubs {
		startDateStr := ymdtoym(elem.StartDate)
		endDateStr := ymdtoym(elem.EndDate)
		allSubsResp = append(allSubsResp,
			models.SubsResponse{
				ServiceName: elem.ServiceName,
				Price:       elem.Price,
				UserID:      elem.UserID,
				StartDate:   startDateStr,
				EndDate:     endDateStr,
			},
		)
	}
	return allSubsResp, nil
}

func (fs *subsServ) UpdateSubs(subsJSON *models.SubsRequest) error {

	if err := validUserID(subsJSON.UserID); err != nil {
		return err
	}

	exists, err := fs.db.CheckUserExists(subsJSON.UserID)
	if err != nil {
		fs.logger.Error().
			Str("error", err.Error()).
			Msg("problem with database")
		return err
	}

	if err := validServiceName(subsJSON.ServiceName); err != nil {
		return err
	}

	if !exists {
		return repository.ErrUserDoesntExists
	}

	err = validPrice(subsJSON.Price)
	if err != nil {
		return err
	}

	startDate, err := ymtoymd(subsJSON.StartDate)
	if err != nil {
		return fmt.Errorf("bad start_date")
	}

	var endDate string
	if subsJSON.EndDate == "" {
		endDate = time.Date(9999, 12, 1, 0, 0, 0, 0, time.UTC).Format("2006-01-02")
	} else {
		endDate, err = ymtoymd(subsJSON.EndDate)
		if err != nil {
			return fmt.Errorf("bad end_date")
		}
	}

	subs := models.Subs{
		ServiceName: subsJSON.ServiceName,
		Price:       subsJSON.Price,
		UserID:      subsJSON.UserID,
		StartDate:   startDate,
		EndDate:     endDate,
	}

	if err := fs.db.UpdateSubs(subs); err != nil {
		fs.logger.Error().
			Str("error", err.Error()).
			Msg("problem with database")
		return err
	}

	return nil
}
