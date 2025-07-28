package handlers

import "github.com/gin-gonic/gin"

type SubsHandler interface {
	CreateSubs(g *gin.Context)
	GetSubs(g *gin.Context)
	DeleteSubs(g *gin.Context)
	GetAllSubs(g *gin.Context)
	GetSumOfSubs(g *gin.Context)
	UpdateSubs(g *gin.Context)
}
