package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// TODO: написать, чтобы каждые 10 секунд пинговалась база данныхх

func createAll(conn *pgx.Conn) error {
	var err error

	if err = сreateSchema(conn); err != nil {
		return fmt.Errorf("can't create schema: %w", err)
	}

	if err = сreateTable(conn); err != nil {
		return fmt.Errorf("can't create table: %w", err)
	}

	return nil
}

// TODO: write this func, this is important!!!
func сreateSchema(conn *pgx.Conn) error {
	query :=
		`
	CREATE SCHEMA IF NOT EXISTS subscriptions
	`
	if _, err := conn.Exec(context.Background(), query); err != nil {
		return fmt.Errorf("can't create scheme: %v", err)
	}

	return nil
}
func сreateTable(conn *pgx.Conn) error {
	query :=
		`
	CREATE TABLE IF NOT EXISTS subscriptions.info (
    	service_name TEXT NOT NULL,
    	price INTEGER NOT NULL,
    	user_id UUID UNIQUE NOT NULL,
    	start_date DATE NOT NULL,
    	end_date DATE NOT NULL 
	)
	`
	if _, err := conn.Exec(context.Background(), query); err != nil {
		return fmt.Errorf("can't create table: %v", err)
	}
	return nil
}
