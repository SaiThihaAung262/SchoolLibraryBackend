package repository

import (
	"MyGO.com/m/model"
	"gorm.io/gorm"
)

type MediaRepository interface {
	CreateMedia(media *model.Media) (*model.Media, error)
}

type mediaConnection struct {
	connection *gorm.DB
}

func NewMediaRepository(connection *gorm.DB) MediaRepository {
	return &mediaConnection{
		connection: connection,
	}
}

func (db *mediaConnection) CreateMedia(media *model.Media) (*model.Media, error) {
	myDb := db.connection.Model(&model.Media{})
	if err := myDb.Create(media).Error; err != nil {
		return nil, err
	}
	return media, nil
}
