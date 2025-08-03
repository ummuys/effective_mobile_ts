package models

import "time"

type SubsRequest struct {
	ServiceName string `json:"service_name" example:"service"`
	Price       int    `json:"price" example:"1"`
	UserID      string `json:"user_id" example:"uuid"`
	StartDate   string `json:"start_date" example:"07-2025"`
	EndDate     string `json:"end_date" example:"optional"`
}

type Subs struct {
	ServiceName string
	Price       int
	UserID      string
	StartDate   string
	EndDate     string
}

type SubsDB struct {
	ServiceName string
	Price       int
	UserID      string
	StartDate   time.Time
	EndDate     time.Time
}

type SubsResponse struct {
	ServiceName string `json:"service_name" example:"Spotify"`
	Price       int    `json:"price" example:"299"`
	UserID      string `json:"user_id" example:"123e4567-e89b-12d3-a456-426614174000"`
	StartDate   string `json:"start_date" example:"07-2025"`
	EndDate     string `json:"end_date" example:"07-2025"`
}

type ErrorResponse struct {
	Message string `json:"msg"`
}
type GoodResponse struct {
	Message string `json:"msg"`
}
