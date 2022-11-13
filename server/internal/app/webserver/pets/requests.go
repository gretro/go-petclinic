package pets

import "petclinic-go/server/internal/app/models"

type CreatePetRequest struct {
	Name      string         `json:"name" binding:"required"`
	BirthDate string         `json:"birthDate" binding:"required"`
	Type      models.PetType `json:"type" binding:"required,oneof=dog cat"`
}

type UpdatePetRequest = CreatePetRequest
