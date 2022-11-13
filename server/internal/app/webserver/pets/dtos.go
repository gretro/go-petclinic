package webserver

import "petclinic-go/server/internal/app/models"

type PetDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`

	BirthDate string         `json:"birthDate"`
	Type      models.PetType `json:"type"`

	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type CreatePetRequest struct {
	Name      string         `json:"name" binding:"required"`
	BirthDate string         `json:"birthDate" binding:"required"`
	Type      models.PetType `json:"type" binding:"required,oneof=dog cat"`
}
