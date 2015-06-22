package models

import "time"

type User struct {
	ID uint `gorm:"primary_key"`
	Name string
	CreatedAt time.Time
	UpdatedAt time.Time
}
