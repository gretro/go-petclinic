package models

import "time"

type Owner struct {
	ID        string
	FirstName string
	LastName  string

	Address     string
	City        string
	PhoneNumber string

	CreatedAt time.Time
	UpdatedAt time.Time

	Pets map[string]*Pet
}
