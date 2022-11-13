package webserver

import (
	"petclinic-go/server/internal/app/models"
	"time"
)

func mapPetToDTO(pet *models.Pet) *PetDTO {
	return &PetDTO{
		ID:   pet.ID,
		Name: pet.Name,

		BirthDate: pet.BirthDate,
		Type:      pet.Type,

		CreatedAt: pet.CreatedAt.UTC().Format(time.RFC3339),
		UpdatedAt: pet.UpdatedAt.UTC().Format(time.RFC3339),
	}
}
