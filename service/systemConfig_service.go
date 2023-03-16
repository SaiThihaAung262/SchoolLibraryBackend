package service

import (
	"fmt"

	"MyGO.com/m/dto"
	"MyGO.com/m/model"
	"MyGO.com/m/repository"
	"github.com/mashingan/smapping"
)

type SystemConfigService interface {
	InsertSystemConfig(systemConfig dto.CreateSystemConfigDto) (*model.SystemConfig, error)
	GetSystemConfig() (*model.SystemConfig, error)
	UpdateSystemConfig(systemConfig dto.UpdateSystemConfigDto) (*model.SystemConfig, error)
	DeleteSystemConfig(id uint64) error
}

type systemConfigService struct {
	systemConfigRepo repository.SystemConfigRepository
}

func NewSystemConfigService(systemConfigRepo repository.SystemConfigRepository) SystemConfigService {
	return &systemConfigService{
		systemConfigRepo: systemConfigRepo,
	}
}

func (service systemConfigService) InsertSystemConfig(systemConfig dto.CreateSystemConfigDto) (*model.SystemConfig, error) {

	var systemConfigToCreate model.SystemConfig
	if err := smapping.FillStruct(&systemConfigToCreate, smapping.MapFields(&systemConfig)); err != nil {
		fmt.Println("--------Here is error in repository ------", err)
	}
	res, err := service.systemConfigRepo.InsertSystemConfig(systemConfigToCreate)
	if err != nil {
		fmt.Println("----------Here is error in update service----------", err)
		return nil, err
	}
	return res, nil
}

func (service systemConfigService) GetSystemConfig() (*model.SystemConfig, error) {
	config, err := service.systemConfigRepo.GetSystemConfig()
	if err != nil {
		return nil, err
	}
	return config, err
}

func (service systemConfigService) UpdateSystemConfig(systemConfig dto.UpdateSystemConfigDto) (*model.SystemConfig, error) {
	configToUpdate := model.SystemConfig{}
	err := smapping.FillStruct(&configToUpdate, smapping.MapFields(&systemConfig))
	if err != nil {
		fmt.Println("------Have error in update bookcategory servcie ------", err.Error())
	}

	res, errRepo := service.systemConfigRepo.UpdateSystemConfig(configToUpdate)
	if errRepo != nil {
		return nil, errRepo
	}

	return res, nil
}

func (service systemConfigService) DeleteSystemConfig(id uint64) error {
	return service.systemConfigRepo.DeleteSystemConfig(id)
}
