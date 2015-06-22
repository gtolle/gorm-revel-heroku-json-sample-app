package models

import "time"

type RecurringVisit struct {
	ID uint `gorm:"primary_key"`
	Frequency int
	CreatedAt time.Time
	UpdatedAt time.Time
}
