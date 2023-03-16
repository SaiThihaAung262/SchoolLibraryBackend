package repository

import (
	"fmt"

	"MyGO.com/m/model"
	"gorm.io/gorm"
)

type PunishmentRepository interface {
	InsertPunishment(punishment model.Punishment) (*model.Punishment, error)
	GetPunishmentData() ([]model.Punishment, error)
	UpdatePunishment(punishment model.Punishment) (*model.Punishment, error)
	DeletePunishment(id uint64) error
}

type punishmentConnection struct {
	connection *gorm.DB
}

func NewPunishmentRepository(db *gorm.DB) PunishmentRepository {
	return &punishmentConnection{
		connection: db,
	}
}

func (db *punishmentConnection) GetPunishmentData() ([]model.Punishment, error) {
	var punishmentData []model.Punishment
	sql := "select * from punishments where deleted_at IS NULL"
	res := db.connection.Raw(sql).Scan(&punishmentData)
	if res.Error != nil {
		return nil, res.Error
	}

	return punishmentData, nil
}

func (db *punishmentConnection) InsertPunishment(punishment model.Punishment) (*model.Punishment, error) {
	err := db.connection.Save(&punishment).Error
	if err != nil {
		fmt.Println("------------Here is error in student repository--------------", err)
		return nil, err
	}
	return &punishment, nil
}

func (db *punishmentConnection) UpdatePunishment(punishment model.Punishment) (*model.Punishment, error) {
	// err := db.connection.Model(&punishment).Where("id = ?", punishment.ID).Updates(model.Punishment{}).Error
	// if err != nil {
	// 	fmt.Println("Error at update student category repository----")
	// 	return nil, err
	// }
	// return &punishment, nil

	updatePunishment := model.Punishment{}

	err := db.connection.Model(&updatePunishment).Where("id = ?", punishment.ID).Select("*").Updates(model.Punishment{
		PackageName:  punishment.PackageName,
		Duration:     punishment.Duration,
		PunishAmount: punishment.PunishAmount,
	}).Error

	if err != nil {
		fmt.Println("----Here have error in update borrow qty book repo -----")
		return nil, err
	}
	return &updatePunishment, nil
}

func (db *punishmentConnection) DeletePunishment(id uint64) error {

	mydb := db.connection.Model(&model.Punishment{})
	mydb = mydb.Where(fmt.Sprintf("id  = %d", id))
	if err := mydb.Delete(&model.Punishment{}).Error; err != nil {
		return err
	}
	return nil
}
