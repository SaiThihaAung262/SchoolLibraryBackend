package service

import (
	"MyGO.com/m/model"
	"MyGO.com/m/repository"
)

type MedeiaService interface {
	CreateMedia(media *model.Media) (*model.Media, error)
}

type mediaService struct {
	mediaRepo repository.MediaRepository
}

func NewMediaService(mediaRepo repository.MediaRepository) MedeiaService {
	return &mediaService{
		mediaRepo: mediaRepo,
	}
}

func (service mediaService) CreateMedia(media *model.Media) (*model.Media, error) {
	res, err := service.mediaRepo.CreateMedia(media)
	return res, err
}
