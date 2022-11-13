package services

import (
	"petclinic-go/server/internal/app/models"
	"petclinic-go/server/internal/app/repository"
	"time"

	"github.com/google/uuid"
)

type VisitsService struct {
	ownersRepository *repository.OwnersRepository
	visitsRepository *repository.VisitsRepository
}

func NewVisitsService(ownersRepository *repository.OwnersRepository, visitsRepository *repository.VisitsRepository) *VisitsService {
	return &VisitsService{
		ownersRepository: ownersRepository,
		visitsRepository: visitsRepository,
	}
}

func (service *VisitsService) GetAll() ([]*models.Visit, error) {
	visits := service.visitsRepository.ListAll(nil)
	return visits, nil
}

func (service *VisitsService) GetAllForPet(petID string) ([]*models.Visit, error) {
	visits := service.visitsRepository.ListAll(&repository.VisitFilters{
		PetID: &petID,
	})

	return visits, nil
}

func (service *VisitsService) GetByID(visitID string) (*models.Visit, error) {
	visit, err := service.visitsRepository.FindByID(visitID)
	if err != nil {
		return nil, err
	}

	return visit, nil
}

type CreateOrUpdateVisitParams struct {
	Date        time.Time
	Description string
	OwnerID     string
	PetID       string
}

func (service *VisitsService) Create(params *CreateOrUpdateVisitParams) (*models.Visit, error) {
	visit, err := service.visitsRepository.Insert(&models.Visit{
		ID:          uuid.NewString(),
		Date:        params.Date,
		Description: params.Description,

		OwnerID: params.OwnerID,
		PetID:   params.PetID,

		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return nil, err
	}

	return visit, nil
}

func (service *VisitsService) Update(visitID string, params *CreateOrUpdateVisitParams) (*models.Visit, error) {
	visit, err := service.visitsRepository.FindByID(visitID)
	if err != nil {
		return nil, err
	}

	updatedVisit, err := service.visitsRepository.Replace(&models.Visit{
		ID:          visitID,
		Date:        params.Date,
		Description: params.Description,

		PetID: params.PetID,

		CreatedAt: visit.CreatedAt,
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return nil, err
	}

	return updatedVisit, nil
}

func (service *VisitsService) Delete(visitID string) error {
	err := service.visitsRepository.DeleteByID(visitID)
	return err
}
