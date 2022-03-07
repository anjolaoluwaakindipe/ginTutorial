package routes

import (
	"github.com/anjolaoluwaakindipe/gintutorial/controllers"
	"github.com/anjolaoluwaakindipe/gintutorial/services"
	"github.com/gin-gonic/gin"
)

var (
	videoService    services.VideoService       = services.New()
	videoController controllers.VideoController = controllers.New(videoService)
)

func videoRoute(superRoute *gin.RouterGroup) {
	videoRouter := superRoute.Group("/posts")
	videoRouter.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	videoRouter.POST("/", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})
}
