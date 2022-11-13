package services

import (
	"petclinic-go/server/internal/app/models"
	"petclinic-go/server/internal/app/repository"
	"time"

	"github.com/google/uuid"
)

type OwnersService struct {
	ownersRepository *repository.OwnersRepository
}

func NewOwnersService(ownersRepository *repository.OwnersRepository) *OwnersService {
	service := &OwnersService{
		ownersRepository: ownersRepository,
	}

	return service
}

func (service *OwnersService) GetByID(ID string) (*models.Owner, error) {
	return service.ownersRepository.FindByID(ID)
}

func (service *OwnersService) GetAll() []*models.Owner {
	return service.ownersRepository.ListOwners()
}

type CreateOrUpdateOwnerParams struct {
	FirstName string
	LastName  string

	Address     string
	City        string
	PhoneNumber string
}

func (service *OwnersService) CreateOwner(params CreateOrUpdateOwnerParams) (*models.Owner, error) {
	ID := uuid.NewString()

	owner := &models.Owner{
		ID:        ID,
		FirstName: params.FirstName,
		LastName:  params.LastName,

		Address:     params.Address,
		City:        params.City,
		PhoneNumber: params.City,

		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),

		Pets: make(map[string]*models.Pet),
	}

	_, error := service.ownersRepository.InsertOwner(owner)
	if error != nil {
		return nil, error
	}

	return owner, nil
}

func (service *OwnersService) UpdateOwner(ownerID string, params CreateOrUpdateOwnerParams) (*models.Owner, error) {
	owner, err := service.ownersRepository.FindByID(ownerID)
	if err != nil {
		return nil, err
	}

	updatedOwner := &models.Owner{
		ID:        ownerID,
		FirstName: params.FirstName,
		LastName:  params.LastName,

		Address:     params.Address,
		City:        params.City,
		PhoneNumber: params.PhoneNumber,

		CreatedAt: owner.CreatedAt,
		UpdatedAt: time.Now(),

		Pets: owner.Pets,
	}

	_, err = service.ownersRepository.ReplaceOwner(updatedOwner)
	if err != nil {
		return nil, err
	}

	return updatedOwner, nil
}

func (service *OwnersService) DeleteOwnerByID(ID string) error {
	return service.ownersRepository.DeleteByID(ID)
}
