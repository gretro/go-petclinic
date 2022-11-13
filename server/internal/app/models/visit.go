package models

import "time"

type Visit struct {
	ID          string
	Date        time.Time
	Description string

	OwnerID string
	PetID   string

	CreatedAt time.Time
	UpdatedAt time.Time
}
