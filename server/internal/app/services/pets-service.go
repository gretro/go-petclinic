package services

import (
	"petclinic-go/server/internal/app/models"
	"petclinic-go/server/internal/app/repository"
	"petclinic-go/server/internal/app/system"
	"time"

	"github.com/google/uuid"
)

type PetsService struct {
	ownersRepository *repository.OwnersRepository
}

func NewPetsService(ownersRepository *repository.OwnersRepository) *PetsService {
	return &PetsService{
		ownersRepository: ownersRepository,
	}
}

type CreateOrUpdatePetParams struct {
	Name string

	BirthDate string
	Type      models.PetType
}

func (service *PetsService) GetPetByID(ownerID string, petID string) (*models.Pet, error) {
	owner, error := service.ownersRepository.FindByID(ownerID)
	if error != nil {
		return nil, error
	}

	pet, exists := owner.Pets[petID]
	if !exists {
		return nil, &system.EntityNotFoundError{
			EntityName: "Pet",
			ID:         petID,
			OtherParams: map[string]interface{}{
				"ownerID": ownerID,
			},
		}
	}

	return pet, nil
}

func (service *PetsService) CreatePet(ownerID string, params CreateOrUpdatePetParams) (*models.Pet, error) {
	owner, error := service.ownersRepository.FindByID(ownerID)
	if error != nil {
		return nil, error
	}

	pet := &models.Pet{
		ID:   uuid.NewString(),
		Name: params.Name,

		BirthDate: params.BirthDate,
		Type:      params.Type,

		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),

		Visits: make(map[string]string),
	}

	owner.Pets[pet.ID] = pet
	return pet, nil
}

func (service *PetsService) GetPetsForOwnerID(ownerID string) ([]*models.Pet, error) {
	owner, error := service.ownersRepository.FindByID(ownerID)
	if error != nil {
		return nil, error
	}

	pets := make([]*models.Pet, len(owner.Pets))

	i := 0
	for _, pet := range owner.Pets {
		pets[i] = pet
		i++
	}

	return pets, nil
}

func (service *PetsService) UpdatePet(ownerID string, petID string, updates *CreateOrUpdatePetParams) (*models.Pet, error) {
	owner, error := service.ownersRepository.FindByID(ownerID)
	if error != nil {
		return nil, error
	}

	pet, exists := owner.Pets[petID]
	if !exists {
		return nil, &system.EntityNotFoundError{
			EntityName: "Pet",
			ID:         petID,
			OtherParams: map[string]interface{}{
				"ownerID": ownerID,
			},
		}
	}

	updatedPet := &models.Pet{
		ID:        petID,
		Name:      updates.Name,
		Type:      updates.Type,
		BirthDate: updates.BirthDate,

		CreatedAt: pet.CreatedAt,
		UpdatedAt: time.Now(),

		Visits: pet.Visits,
	}

	owner.Pets[petID] = updatedPet

	return updatedPet, nil
}

func (service *PetsService) DeletePet(ownerID string, petID string) error {
	owner, error := service.ownersRepository.FindByID(ownerID)
	if error != nil {
		return error
	}

	delete(owner.Pets, petID)
	return nil
}
