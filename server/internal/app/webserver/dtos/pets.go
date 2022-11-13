package dtos

import (
	"petclinic-go/server/internal/app/models"
	"time"
)

type PetSummaryDTO struct {
	ID   string         `json:"id"`
	Name string         `json:"name"`
	Type models.PetType `json:"type"`
}

type PetDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`

	BirthDate string         `json:"birthDate"`
	Type      models.PetType `json:"type"`

	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func MapPetToDTO(pet *models.Pet) *PetDTO {
	return &PetDTO{
		ID:   pet.ID,
		Name: pet.Name,

		BirthDate: pet.BirthDate,
		Type:      pet.Type,

		CreatedAt: pet.CreatedAt.UTC().Format(time.RFC3339),
		UpdatedAt: pet.UpdatedAt.UTC().Format(time.RFC3339),
	}
}

func MapPetToSummaryDTO(pet *models.Pet) *PetSummaryDTO {
	if pet == nil {
		return nil
	}

	return &PetSummaryDTO{
		ID:   pet.ID,
		Name: pet.Name,
		Type: pet.Type,
	}
}
