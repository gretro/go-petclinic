package dtos

import (
	"petclinic-go/server/internal/app/models"
	"time"
)

type OwnerDTO struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`

	Address     string `json:"address"`
	City        string `json:"city"`
	PhoneNumber string `json:"phoneNumber"`

	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type OwnerSummaryDTO struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`

	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func MapOwnerToDTO(owner *models.Owner) *OwnerDTO {
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

func MapOwnerToSummaryDTO(owner *models.Owner) *OwnerSummaryDTO {
	if owner == nil {
		return nil
	}

	return &OwnerSummaryDTO{
		ID:        owner.ID,
		FirstName: owner.FirstName,
		LastName:  owner.LastName,

		CreatedAt: owner.CreatedAt.UTC().Format(time.RFC3339),
		UpdatedAt: owner.UpdatedAt.UTC().Format(time.RFC3339),
	}
}
