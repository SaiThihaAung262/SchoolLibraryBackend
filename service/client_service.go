package service

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"MyGO.com/m/dto"
	"MyGO.com/m/model"
	"MyGO.com/m/repository"
	"github.com/mashingan/smapping"
)

type ClientService interface {
	InsertClient(client dto.ClientRegisterDTO) (*model.Client, error)
	GetAllClients(req *dto.ClientGetRequest) ([]model.Client, int64, error)
	UpdateClient(client dto.UpdateClientDTO) (*model.Client, error)
	DeleteClient(id uint64) error
}
type clientService struct {
	clientRepo repository.ClientRepository
}

func NewClientService(clientRepo repository.ClientRepository) ClientService {
	return &clientService{
		clientRepo: clientRepo,
	}
}

func GenerateUUID() string {
	newUUID, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal(err)
	}
	splitID := strings.Split(string(newUUID), "-")
	fmt.Println("Generated UUID:")
	fmt.Printf("%s", splitID[0])
	return splitID[0]
}

func (service clientService) InsertClient(client dto.ClientRegisterDTO) (*model.Client, error) {
	var clientToCreate model.Client
	if err := smapping.FillStruct(&clientToCreate, smapping.MapFields(&client)); err != nil {
		fmt.Println("--------Here is error in repository ------", err)
	}
	clientToCreate.UUID = GenerateUUID()
	res, err := service.clientRepo.InsertClient(clientToCreate)
	if err != nil {
		fmt.Println("----------Here is error in update service----------", err)
		return nil, err
	}
	return res, nil
}

func (service clientService) GetAllClients(req *dto.ClientGetRequest) ([]model.Client, int64, error) {

	clients, count, err := service.clientRepo.GetAllClients(req)
	if err != nil {
		return nil, 0, err
	}
	return clients, count, err
}

func (service clientService) UpdateClient(client dto.UpdateClientDTO) (*model.Client, error) {
	clientToUpdate := model.Client{}
	err := smapping.FillStruct(&clientToUpdate, smapping.MapFields(client))
	if err != nil {
		fmt.Println("------Have error in update bookcategory servcie ------", err.Error())
	}

	res, errRepo := service.clientRepo.UpdateClient(clientToUpdate)
	if errRepo != nil {
		return nil, errRepo
	}

	return res, nil

}

func (service clientService) DeleteClient(id uint64) error {
	err := service.clientRepo.DeleteClient(id)
	return err
}
