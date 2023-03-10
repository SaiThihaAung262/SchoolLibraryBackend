package repository

import (
	"fmt"

	"MyGO.com/m/dto"
	"MyGO.com/m/model"
	"gorm.io/gorm"
)

type ClientRepository interface {
	InsertClient(client model.Client) (*model.Client, error)
	GetAllClients(req *dto.ClientGetRequest) ([]model.Client, int64, error)
	UpdateClient(client model.Client) (*model.Client, error)
	DeleteClient(id uint64) error
}

type clientConnection struct {
	connection *gorm.DB
}

func NewClientRepository(db *gorm.DB) ClientRepository {
	return &clientConnection{
		connection: db,
	}
}

func (db *clientConnection) InsertClient(client model.Client) (*model.Client, error) {
	err := db.connection.Save(&client).Error
	if err != nil {
		fmt.Println("------------Here is error in user repository--------------", err)
		return nil, err
	}
	return &client, nil
}

func (db *clientConnection) GetAllClients(req *dto.ClientGetRequest) ([]model.Client, int64, error) {
	var clients []model.Client
	var total int64

	var offset uint64
	var pageSize uint64

	if req.Page != 0 {
		offset = (req.Page - 1) * req.PageSize
	} else {
		offset = 0
	}

	if req.PageSize != 0 {
		pageSize = req.PageSize
	} else {
		pageSize = 10

	}

	filter := " where deleted_at IS NULL"

	if req.UUID != "" {
		filter += fmt.Sprintf(" and uuid = %s", req.UUID)
	}

	if req.Type != 0 {
		filter += fmt.Sprintf(" and type = %d", req.Type)
	}

	if req.Name != "" {
		filter += fmt.Sprintf(" and name LIKE \"%s%s%s\"", "%", req.Name, "%")
	}

	if req.Email != "" {
		filter += fmt.Sprintf(" and email LIKE \"%s%s%s\"", "%", req.Email, "%")

	}

	countQuery := fmt.Sprintf("select count(1) from clients %s", filter)
	if err := db.connection.Raw(countQuery).Scan(&total).Error; err != nil {
		return nil, 0, err
	}

	sql := fmt.Sprintf("select * from clients %s limit %v offset %v", filter, pageSize, offset)
	res := db.connection.Raw(sql).Scan(&clients)
	if res.Error != nil {
		return nil, 0, res.Error
	}
	return clients, total, nil
}

func (db *clientConnection) UpdateClient(client model.Client) (*model.Client, error) {
	err := db.connection.Model(&client).Where("id = ?", client.ID).Updates(model.Client{
		UUID:     client.UUID,
		Name:     client.Name,
		Email:    client.Email,
		Password: client.Password,
		Type:     client.Type,
	}).Error
	if err != nil {
		fmt.Println("Error at update book category repository----")
		return nil, err
	}
	return &client, nil
}

func (db *clientConnection) DeleteClient(id uint64) error {

	mydb := db.connection.Model(&model.Client{})
	mydb = mydb.Where(fmt.Sprintf("id  = %d", id))
	if err := mydb.Delete(&model.Client{}).Error; err != nil {
		return err
	}
	return nil
}
