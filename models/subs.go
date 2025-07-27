package models

import "time"

type CreateSubsRequest struct {
	ServiceName string `json:"service_name"`
	Price       string `json:"price"`
	UserID      string `json:"user_id"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
}

type Subs struct {
	ServiceName string
	Price       int
	UserID      string
	StartDate   string
	EndDate     string
}

type GetSubs struct {
	ServiceName string
	Price       int
	UserID      string
	StartDate   time.Time
	EndDate     time.Time
}

type GetSubsResponse struct {
	ServiceName string `json:"service_name"`
	Price       string `json:"price"`
	UserID      string `json:"user_id"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
}
