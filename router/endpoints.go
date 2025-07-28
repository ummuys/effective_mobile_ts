package router

// MAIN WAY
const (
	Mainway = "http://127.0.0.1"
)

// SERVER-STATUS
const (
	Health = "/health"
)

// TASKS-API
const (
	CreateSubsWay = "/api/v1/create-subs"
	UpdateSubsWay = "/api/v1/update-subs/:user_id"
	GetSubsWay    = "/api/v1/get-subs/:user_id"
	DeleteSubsWay = "/api/v1/delete-subs/:user_id"
	GetAllSubsWay = "/api/v1/get-subs"
	GetSumOfSubs  = "/api/v1/get-sum-subs"
)
