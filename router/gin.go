package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/ummuys/effective_mobile_ts/docs"
	handlers "github.com/ummuys/effective_mobile_ts/handlers/subscription"
)

func RunRouter(subsHandler handlers.SubsHandler) {

	g := gin.New()
	g.Use(gin.Recovery())

	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	g.POST(CreateSubsWay, subsHandler.CreateSubs)   //    C
	g.GET(GetSubsWay, subsHandler.GetSubs)          // 	R
	g.PUT(UpdateSubsWay, subsHandler.UpdateSubs)    //    U
	g.DELETE(DeleteSubsWay, subsHandler.DeleteSubs) // 	D
	g.GET(GetAllSubsWay, subsHandler.GetAllSubs)    // 	L

	g.Run(":8082")
}
