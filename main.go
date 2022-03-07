package main

import (
	"github.com/anjolaoluwaakindipe/gintutorial/controllers"
	"github.com/anjolaoluwaakindipe/gintutorial/middlewares"
	"github.com/anjolaoluwaakindipe/gintutorial/routes"
	"github.com/anjolaoluwaakindipe/gintutorial/services"
	"github.com/gin-gonic/gin"
)

var (
	videoService    services.VideoService       = services.New()
	videoController controllers.VideoController = controllers.New(videoService)
)

func main() {
	server := gin.New()
	server.Use(gin.Recovery(), middlewares.Logger())

	router := server.Group("/api/v1")

	routes.AddRoutes(router)

	// server.GET("/posts", func(ctx *gin.Context) {
	// 	ctx.JSON(200, videoController.FindAll())
	// })

	// server.POST("/posts", func(ctx *gin.Context){
	// 	ctx.JSON(200, videoController.Save(ctx))
	// })

	server.Run(":8080")

}
