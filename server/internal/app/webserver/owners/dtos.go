package webserver

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

type CreateAndUpdateOwnerRequest struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`

	Address     string `json:"address" binding:"required"`
	City        string `json:"city" binding:"required"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
}

type UpdateOwnerUriParams struct {
	OwnerID string `uri:"ownerId" binding:"required,uuid"`
}
