package repository

import (
	"fmt"

	"MyGO.com/m/model"
	"gorm.io/gorm"
)

type SystemConfigRepository interface {
	InsertSystemConfig(systemConfig model.SystemConfig) (*model.SystemConfig, error)
	GetSystemConfig() (*model.SystemConfig, error)
	UpdateSystemConfig(systemConfig model.SystemConfig) (*model.SystemConfig, error)
	DeleteSystemConfig(id uint64) error
}

type systemConfigConnection struct {
	connection *gorm.DB
}

func NewSystemConfigRepo(db *gorm.DB) SystemConfigRepository {
	return &systemConfigConnection{
		connection: db,
	}
}

func (db *systemConfigConnection) InsertSystemConfig(systemConfig model.SystemConfig) (*model.SystemConfig, error) {
	err := db.connection.Save(&systemConfig).Error
	if err != nil {
		fmt.Println("------------Here is error in student repository--------------", err)
		return nil, err
	}
	return &systemConfig, nil
}

func (db *systemConfigConnection) GetSystemConfig() (*model.SystemConfig, error) {

	var systemConfig model.SystemConfig
	sql := "SELECT * FROM system_configs LIMIT 1;"
	res := db.connection.Raw(sql).Scan(&systemConfig)
	if res.Error != nil {
		return nil, res.Error
	}

	return &systemConfig, nil

}

func (db *systemConfigConnection) UpdateSystemConfig(systemConfig model.SystemConfig) (*model.SystemConfig, error) {

	updateSystemConfig := model.SystemConfig{}

	err := db.connection.Model(&updateSystemConfig).Where("id = ?", systemConfig.ID).Select("*").Updates(model.SystemConfig{
		TeacherCanBorrowCount: systemConfig.TeacherCanBorrowCount,
		StudentCanBorrowCount: systemConfig.StudentCanBorrowCount,
	}).Error

	if err != nil {
		fmt.Println("----Here have error in update borrow qty book repo -----")
		return nil, err
	}
	return &updateSystemConfig, nil

}

func (db *systemConfigConnection) DeleteSystemConfig(id uint64) error {
	mydb := db.connection.Model(&model.SystemConfig{})
	mydb = mydb.Where(fmt.Sprintf("id  = %d", id))
	if err := mydb.Delete(&model.SystemConfig{}).Error; err != nil {
		return err
	}
	return nil
}
