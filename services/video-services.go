package services

import "github.com/anjolaoluwaakindipe/gintutorial/entities"

type VideoService interface {
	Save(entities.Video) entities.Video;
	FindAll() []entities.Video
}

type videoService struct{
	videos []entities.Video
}

func New() *videoService{
	return &videoService{}
}

func (service *videoService) Save(video entities.Video) entities.Video{
	service.videos = append(service.videos, video)
	return video
}

func (service *videoService) FindAll() []entities.Video{
	return service.videos
}