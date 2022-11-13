package visits

type CreateVisitRequest struct {
	Date        string `json:"date" binding:"required"`
	Description string `json:"description"`

	PetID   string `json:"petId" binding:"required,uuid"`
	OwnerID string `json:"ownerId" binding:"required,uuid"`
}

type UpdateVisitRequest = CreateVisitRequest
