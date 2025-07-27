package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/ummuys/effective_mobile_ts/models"
	"github.com/ummuys/effective_mobile_ts/repository"
	service "github.com/ummuys/effective_mobile_ts/service/subscription"
)

type subsHand struct {
	subsService service.SubsService
	logger      *zerolog.Logger
}

func NewSubsHandler(subsService service.SubsService, logger *zerolog.Logger) SubsHandler {
	return &subsHand{
		subsService: subsService,
		logger:      logger,
	}
}

// CreateSubs godoc
// @Summary Создать подписку
// @Description Создаёт новую подписку по данным из тела запроса
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param request body models.CreateSubsRequest true "Данные подписки"
// @Success 200 {object} models.GoodResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/create-subs [post]
func (fh *subsHand) CreateSubs(g *gin.Context) {
	reqJSON := models.CreateSubsRequest{}
	if err := g.ShouldBindJSON(&reqJSON); err != nil {
		fh.logger.Warn().
			Str("client_ip", g.ClientIP()).
			Str("method", "POST").
			Str("path", g.Request.URL.Path).
			Msg("bad request")
		g.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}

	if err := fh.subsService.CreateSubs(&reqJSON); err != nil {
		if errors.Is(err, repository.ErrDBUnavailable) {
			g.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{Message: "something with server, try again later"})
			return
		}
		fh.logger.Warn().
			Str("client_ip", g.ClientIP()).
			Str("method", "POST").
			Str("path", g.Request.URL.Path).
			Msg(err.Error())
		g.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}

	fh.logger.Info().
		Str("client_ip", g.ClientIP()).
		Str("method", "POST").
		Str("path", g.Request.URL.Path).
		Msg("follow created")

	g.JSON(http.StatusOK, models.GoodResponse{Message: "follow created"})
}

// GetSubs godoc
// @Summary Получить подписку по user_id
// @Description Возвращает информацию о подписке по ID пользователя
// @Tags subscriptions
// @Produce json
// @Param user_id path string true "ID пользователя"
// @Success 200 {object} models.GetSubsResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/get-subs/{user_id} [get]
func (fh *subsHand) GetSubs(g *gin.Context) {
	userID := g.Param("user_id")
	resp, err := fh.subsService.GetSubs(userID)
	if err != nil {
		if errors.Is(err, repository.ErrDBUnavailable) {
			g.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{Message: "something with server, try again later"})
			return
		}
		fh.logger.Warn().
			Str("client_ip", g.ClientIP()).
			Str("method", "GET").
			Str("path", g.Request.URL.Path).
			Msg(err.Error())
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	fh.logger.Info().
		Str("client_ip", g.ClientIP()).
		Str("method", "DELETE").
		Str("path", g.Request.URL.Path).
		Msg("returned subs")
	g.JSON(http.StatusOK, resp)
}

// DeleteSubs godoc
// @Summary Удалить подписку по user_id
// @Description Удаляет подписку, связанную с указанным пользователем
// @Tags subscriptions
// @Produce json
// @Param user_id path string true "ID пользователя"
// @Success 200 {object} models.GoodResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/delete-subs/{user_id} [delete]
func (fh *subsHand) DeleteSubs(g *gin.Context) {
	userID := g.Param("user_id")
	err := fh.subsService.DeleteSubs(userID)
	if err != nil {
		if errors.Is(err, repository.ErrDBUnavailable) {
			g.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{Message: "something with server, try again later"})
			return
		}
		fh.logger.Warn().
			Str("client_ip", g.ClientIP()).
			Str("method", "DELETE").
			Str("path", g.Request.URL.Path).
			Msg(err.Error())
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	msg := fmt.Sprintf("subs with user_id = %s deleted", userID)
	fh.logger.Info().
		Str("client_ip", g.ClientIP()).
		Str("method", "DELETE").
		Str("path", g.Request.URL.Path).
		Msg(msg)

	g.JSON(http.StatusOK, gin.H{"msg": msg})
}

// GetAllSubs godoc
// @Summary Получить все подписки
// @Description Возвращает список всех подписок
// @Tags subscriptions
// @Produce json
// @Success 200 {array} models.GetSubsResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/get-subs [get]
func (fh *subsHand) GetAllSubs(g *gin.Context) {
	allSubsResp, err := fh.subsService.GetAllSubs()
	if err != nil {
		if errors.Is(err, repository.ErrDBUnavailable) {
			g.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{Message: "something with server, try again later"})
			return
		}
		fh.logger.Warn().
			Str("client_ip", g.ClientIP()).
			Str("method", "GET").
			Str("path", g.Request.URL.Path).
			Msg(err.Error())
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	if allSubsResp == nil {
		fh.logger.Info().
			Str("client_ip", g.ClientIP()).
			Str("method", "GET").
			Str("path", g.Request.URL.Path).
			Msg("no subscruption exists")
		g.JSON(http.StatusOK, gin.H{"msg": "no subscription exists"})
		return
	}
	msg := fmt.Sprintf("returned %d subs", len(allSubsResp))
	fh.logger.Info().
		Str("client_ip", g.ClientIP()).
		Str("method", "GET").
		Str("path", g.Request.URL.Path).
		Msg(msg)
	g.JSON(http.StatusOK, allSubsResp)
}

// UpdateSubs godoc
// @Summary Обновить подписку
// @Description Обновляет данные подписки по данным из тела запроса
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param request body models.CreateSubsRequest true "Обновлённые данные подписки"
// @Success 200 {array} models.GetSubsResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/update-subs/{user_id} [put]
func (fh *subsHand) UpdateSubs(g *gin.Context) {
	reqJSON := models.CreateSubsRequest{}
	if err := g.ShouldBindJSON(&reqJSON); err != nil {
		fh.logger.Warn().
			Str("client_ip", g.ClientIP()).
			Str("method", "PUT").
			Str("path", g.Request.URL.Path).
			Msg("bad request")
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "request should bind json"})
		return
	}

	if err := fh.subsService.UpdateSubs(&reqJSON); err != nil {
		if errors.Is(err, repository.ErrDBUnavailable) {
			g.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{Message: "something with server, try again later"})
			return
		}
		fh.logger.Warn().
			Str("client_ip", g.ClientIP()).
			Str("method", "PUT").
			Str("path", g.Request.URL.Path).
			Msg(err.Error())
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	fh.logger.Info().
		Str("client_ip", g.ClientIP()).
		Str("method", "PUT").
		Str("path", g.Request.URL.Path).
		Msg("follow updated")

	g.JSON(http.StatusOK, gin.H{"msg": "follow updated"})
}
