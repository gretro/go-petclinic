package repository

import (
	"petclinic-go/server/internal/app/models"
	"petclinic-go/server/internal/app/system"
)

type OwnersRepository struct {
	owners map[string]*models.Owner
}

func NewOwnersRepository() *OwnersRepository {
	return &OwnersRepository{
		owners: make(map[string]*models.Owner),
	}
}

func (repo *OwnersRepository) InsertOwner(owner *models.Owner) (*models.Owner, error) {
	_, exists := repo.owners[owner.ID]
	if exists {
		return nil, &system.EntityConflictError{
			EntityName: "Owner",
			ID:         owner.ID,
		}
	}

	repo.owners[owner.ID] = owner

	return owner, nil
}

func (repo *OwnersRepository) ReplaceOwner(owner *models.Owner) (*models.Owner, error) {
	repo.owners[owner.ID] = owner

	return owner, nil
}

func (repo *OwnersRepository) ListOwners() []*models.Owner {
	owners := make([]*models.Owner, len(repo.owners))

	i := 0
	for _, value := range repo.owners {
		owners[i] = value
		i++
	}

	return owners
}

func (repo *OwnersRepository) FindByID(ID string) (*models.Owner, error) {
	owner, exists := repo.owners[ID]

	if !exists {
		return nil, &system.EntityNotFoundError{
			EntityName:  "Owner",
			ID:          ID,
			OtherParams: make(map[string]interface{}),
		}
	}

	return owner, nil
}

func (repo *OwnersRepository) DeleteByID(ID string) error {
	delete(repo.owners, ID)

	return nil
}
