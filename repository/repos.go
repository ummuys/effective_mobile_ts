package repository

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog"
	"github.com/ummuys/effective_mobile_ts/models"
)

type dbPg struct {
	Conn   *pgx.Conn
	logger *zerolog.Logger
}

func NewDatabase(logger *zerolog.Logger) (Database, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var (
		conn *pgx.Conn
		err  error
	)

	for i := 0; i < 10; i++ {
		conn, err = pgx.Connect(ctx, os.Getenv("DB_LINK"))
		if err == nil {
			break
		}
		time.Sleep(time.Second)
	}
	if err != nil {
		return nil, fmt.Errorf("can't connect to db: %w", err)
	}

	if err = conn.Ping(ctx); err != nil {
		return nil, fmt.Errorf("db didn't pinged: %w", err)
	}

	if err = createAll(conn); err != nil {
		return nil, err
	}

	return &dbPg{
		Conn:   conn,
		logger: logger,
	}, nil

}

func (db *dbPg) CheckUserExists(userID string) (bool, error) {

	if err := db.Conn.Ping(context.Background()); err != nil {
		db.logger.Error().Msg("ping didn't return answer: " + err.Error())
		return false, ErrDBUnavailable
	}

	query := `
	SELECT EXISTS (
		SELECT 1 FROM subscriptions.info
		WHERE user_id = $1
	)
	`
	var exists bool
	if err := db.Conn.QueryRow(context.Background(), query, userID).Scan(&exists); err != nil {
		return false, fmt.Errorf("can't make query: %w", err)
	}
	return exists, nil
}

func (db *dbPg) CreateSubs(subsInfo models.Subs) error {

	if err := db.Conn.Ping(context.Background()); err != nil {
		db.logger.Error().Msg("ping didn't return answer: " + err.Error())
		return ErrDBUnavailable
	}

	query := `
	INSERT INTO subscriptions.info (service_name, price, user_id, start_date, end_date) VALUES
	($1, $2, $3, $4, $5)
	`

	if _, err := db.Conn.Exec(context.Background(),
		query,
		subsInfo.ServiceName,
		subsInfo.Price,
		subsInfo.UserID,
		subsInfo.StartDate,
		subsInfo.EndDate,
	); err != nil {
		return fmt.Errorf("can't make query: %w", err)
	}
	return nil
}

func (db *dbPg) GetSubs(userID string) (*models.SubsDB, error) {

	if err := db.Conn.Ping(context.Background()); err != nil {
		db.logger.Error().Msg("ping didn't return answer: " + err.Error())
		return nil, ErrDBUnavailable
	}

	query := `
	select * from subscriptions.info
	where user_id = $1
	`

	subs := models.SubsDB{}

	err := db.Conn.QueryRow(context.Background(), query, userID).Scan(&subs.ServiceName, &subs.Price, &subs.UserID, &subs.StartDate, &subs.EndDate)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			db.logger.Info().
				Str("user_id", userID).
				Msg("no subscriptions found")
			return nil, ErrUserDoesntExists
		}
		return nil, fmt.Errorf("can't make query: %w", err)
	}
	return &subs, nil
}

func (db *dbPg) DeleteSubs(userID string) error {

	if err := db.Conn.Ping(context.Background()); err != nil {
		db.logger.Error().Msg("ping didn't return answer: " + err.Error())
		return ErrDBUnavailable
	}

	query := `
	DELETE FROM subscriptions.info 
	where user_id = $1
	`
	_, err := db.Conn.Exec(context.Background(), query, userID)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return fmt.Errorf("can't make query: %w", err)
		}
	}
	return nil
}

func (db *dbPg) GetAllSubs() ([]models.SubsDB, error) {

	if err := db.Conn.Ping(context.Background()); err != nil {
		db.logger.Error().Msg("ping didn't return answer: " + err.Error())
		return nil, ErrDBUnavailable
	}
	query := `
	select * from subscriptions.info
	`

	rows, err := db.Conn.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("can't make query: %w", err)
	}
	defer rows.Close()

	var allSubs []models.SubsDB

	for rows.Next() {
		var subs models.SubsDB
		err := rows.Scan(&subs.ServiceName, &subs.Price, &subs.UserID, &subs.StartDate, &subs.EndDate)
		if err != nil {
			return nil, fmt.Errorf("can't fill a data: %w", err)
		}
		allSubs = append(allSubs, subs)
	}

	//TODO: change this err
	if rows.Err() != nil {
		return nil, fmt.Errorf("some conflict: %w", err)
	}

	if len(allSubs) == 0 {
		return nil, nil
	}
	return allSubs, nil
}

func (db *dbPg) UpdateSubs(subsInfo models.Subs) error {

	if err := db.Conn.Ping(context.Background()); err != nil {
		db.logger.Error().Msg("ping didn't return answer: " + err.Error())
		return ErrDBUnavailable
	}

	query := `
	UPDATE subscriptions.info
	SET service_name = $1,
    	price = $2,
    	start_date = $3,
    	end_date = $4
	where user_id = $5 
	`

	_, err := db.Conn.Exec(context.Background(), query, subsInfo.ServiceName, subsInfo.Price, subsInfo.StartDate, subsInfo.EndDate, subsInfo.UserID)
	if err != nil {
		return fmt.Errorf("can't make query: %w", err)
	}

	return nil
}
