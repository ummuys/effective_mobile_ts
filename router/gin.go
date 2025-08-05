package router

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/ummuys/effective_mobile_ts/docs"
	handlers "github.com/ummuys/effective_mobile_ts/handlers/subscription"
)

func CreateServer(subsHandler handlers.SubsHandler) *http.Server {

	g := gin.New()
	g.Use(gin.Recovery())
	gin.SetMode(gin.ReleaseMode)

	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	g.POST(CreateSubsWay, subsHandler.CreateSubs)   //    C
	g.GET(GetSubsWay, subsHandler.GetSubs)          // 	  R
	g.PUT(UpdateSubsWay, subsHandler.UpdateSubs)    //    U
	g.DELETE(DeleteSubsWay, subsHandler.DeleteSubs) // 	  D
	g.GET(GetAllSubsWay, subsHandler.GetAllSubs)    // 	  L

	g.GET(GetSumOfSubs, subsHandler.GetSumOfSubs) // 	  S

	server := &http.Server{
		Addr:    os.Getenv("IP") + ":" + os.Getenv("PORT"),
		Handler: g,
	}

	return server
}

func RunServer(server *http.Server, logger *zerolog.Logger) {
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		msg := fmt.Sprintf("listen error: %v", err)
		logger.Fatal().Msg(msg)
	}
}
