package models

import "time"

type VisitItem struct {
	ID uint `gorm:"primary_key"`
	VisitId uint
	Visit Visit
	ItemId uint
	Quantity uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

type VisitItemJson struct {
	ID uint `json:"id"`
	VisitId uint `json:"visit_id"`
	ItemId uint `json:"item_id"`
	Quantity uint `json:"quantity"`
}

func (visitItem *VisitItem) ToJson() VisitItemJson {
	return VisitItemJson {
		ID: visitItem.ID,
		VisitId: visitItem.VisitId,
		ItemId: visitItem.ItemId,
		Quantity: visitItem.Quantity,
	}
}
