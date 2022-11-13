package repository

import (
	"petclinic-go/server/internal/app/models"
	"petclinic-go/server/internal/app/system"
)

type VisitsRepository struct {
	visits map[string]*models.Visit
}

func NewVisitsRepository() *VisitsRepository {
	return &VisitsRepository{
		visits: make(map[string]*models.Visit),
	}
}

type VisitFilters struct {
	PetID *string
}

func (repo *VisitsRepository) ListAll(filters *VisitFilters) []*models.Visit {
	visits := make([]*models.Visit, 0, len(repo.visits))

	for _, value := range repo.visits {
		if filters != nil {
			if filters.PetID != nil && value.PetID == *filters.PetID {
				visits = append(visits, value)
			}
		} else {
			visits = append(visits, value)
		}
	}

	return visits
}

func (repo *VisitsRepository) Insert(visit *models.Visit) (*models.Visit, error) {
	_, exists := repo.visits[visit.ID]
	if exists {
		return nil, &system.EntityConflictError{
			EntityName: "Visit",
			ID:         visit.ID,
		}
	}

	repo.visits[visit.ID] = visit
	return visit, nil
}

func (repo *VisitsRepository) FindByID(visitID string) (*models.Visit, error) {
	visit, exists := repo.visits[visitID]
	if !exists {
		return nil, &system.EntityNotFoundError{
			EntityName: "Visit",
			ID:         visitID,
		}
	}

	return visit, nil
}

func (repo *VisitsRepository) Replace(visit *models.Visit) (*models.Visit, error) {
	_, exists := repo.visits[visit.ID]
	if !exists {
		return nil, &system.EntityNotFoundError{
			EntityName: "Visit",
			ID:         visit.ID,
		}
	}

	repo.visits[visit.ID] = visit

	return visit, nil
}

func (repo *VisitsRepository) DeleteByID(visitID string) error {
	delete(repo.visits, visitID)

	return nil
}
