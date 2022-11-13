package models

import "time"

type Pet struct {
	ID   string
	Name string

	BirthDate string
	Type      PetType

	CreatedAt time.Time
	UpdatedAt time.Time

	Visits map[string]string
}
