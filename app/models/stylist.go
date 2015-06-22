package models

import "time"

type Stylist struct {
	ID uint `gorm:"primary_key"`
	UserId uint
	User User
	CreatedAt time.Time
	UpdatedAt time.Time
}
