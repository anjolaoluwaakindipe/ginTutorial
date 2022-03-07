package controllers

import (
	"github.com/anjolaoluwaakindipe/gintutorial/entities"
	"github.com/anjolaoluwaakindipe/gintutorial/services"
	"github.com/gin-gonic/gin"
)

type VideoController interface {
	FindAll() []entities.Video
	Save(ctx *gin.Context) entities.Video
}

type controller struct {
	service services.VideoService
}

func New(service services.VideoService) *controller {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entities.Video{
	videos:= c.service.FindAll()
	return videos
}

func (c *controller ) Save(ctx *gin.Context) entities.Video {
	var video entities.Video
	ctx.BindJSON(&video)
	c.service.Save(video);
	return video
}
