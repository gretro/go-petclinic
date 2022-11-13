package dtos

import (
	"petclinic-go/server/internal/app/models"
	"time"
)

type VisitDTO struct {
	ID          string `json:"id"`
	Date        string `json:"date"`
	Description string `json:"description"`

	Pet   *PetSummaryDTO   `json:"pet"`
	Owner *OwnerSummaryDTO `json:"owner"`

	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func MapVisitToDTO(visit *models.Visit, owner *models.Owner, pet *models.Pet) *VisitDTO {
	return &VisitDTO{
		ID:          visit.ID,
		Date:        visit.Date.UTC().Format(time.RFC3339),
		Description: visit.Description,

		Pet:   MapPetToSummaryDTO(pet),
		Owner: MapOwnerToSummaryDTO(owner),

		CreatedAt: visit.CreatedAt.UTC().Format(time.RFC3339),
		UpdatedAt: visit.UpdatedAt.UTC().Format(time.RFC3339),
	}
}
