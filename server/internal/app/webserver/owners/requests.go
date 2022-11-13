package owners

type CreateAndUpdateOwnerRequest struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`

	Address     string `json:"address" binding:"required"`
	City        string `json:"city" binding:"required"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
}
