package models

import "time"

type Visit struct {
	ID          string
	Date        time.Time
	Description string
}
