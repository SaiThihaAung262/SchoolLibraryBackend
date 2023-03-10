package repository

import (
	"fmt"

	"MyGO.com/m/model"
	"gorm.io/gorm"
)

type ClientRepository interface {
	InsertClient(client model.Client) (*model.Client, error)
	// GetAllClients(req *dto.UserGetRequest) ([]model.User, int64, error)
	// UpdateClient(user model.User) (*model.User, error)
	// DeleteClient(id uint64) error
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
