package services

import (
	"petclinic-go/server/internal/app/models"
	"petclinic-go/server/internal/app/repository"
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

type CreatePetParams struct {
	Name string

	BirthDate string
	Type      models.PetType
}

func (service *PetsService) CreatePet(ownerID string, params CreatePetParams) (*models.Pet, error) {
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

		Visits: make(map[string]*models.Visit),
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
