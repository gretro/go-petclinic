package webserver

import (
	"petclinic-go/server/internal/app/models"
	"time"
)

func mapOwnerToDTO(owner *models.Owner) *OwnerDTO {
	return &OwnerDTO{
		ID:        owner.ID,
		FirstName: owner.FirstName,
		LastName:  owner.LastName,

		Address:     owner.Address,
		City:        owner.City,
		PhoneNumber: owner.PhoneNumber,

		CreatedAt: owner.CreatedAt.UTC().Format(time.RFC3339),
		UpdatedAt: owner.UpdatedAt.UTC().Format(time.RFC3339),
	}
}

func mapOwnerToSummaryDTO(owner *models.Owner) *OwnerSummaryDTO {
	return &OwnerSummaryDTO{
		ID:        owner.ID,
		FirstName: owner.FirstName,
		LastName:  owner.LastName,

		CreatedAt: owner.CreatedAt.UTC().Format(time.RFC3339),
		UpdatedAt: owner.UpdatedAt.UTC().Format(time.RFC3339),
	}
}
