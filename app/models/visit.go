package models

import "time"

type Visit struct {
	ID uint `gorm:"primary_key"`
	UserId uint
	User User
	StylistId uint
	Stylist Stylist
	Date time.Time
	RecurringVisitId uint
	RecurringVisit RecurringVisit
	VisitState string
	VisitItems []VisitItem
	CreatedAt time.Time
	UpdatedAt time.Time
}

type VisitJson struct {
	ID uint `json:"id"`
	UserId uint `json:"user_id"`
	UserName string `json:"user_name"`
	Date time.Time `json:"date"`
	StylistName string `json:"stylist_name"`
	IsRecurring bool `json:"is_recurring"`
	RecurrenceFrequency int `json:"recurrence_frequency"`
	VisitItems []VisitItemJson `json:"visit_items,omitempty"`
}

func (visit *Visit) ToJson() VisitJson {
	return VisitJson {
		ID: visit.ID,
		UserId: visit.UserId,
		UserName: visit.User.Name,
		Date: visit.Date,
		StylistName: visit.Stylist.User.Name,
		IsRecurring: visit.RecurringVisit.ID != 0,
		RecurrenceFrequency: visit.RecurringVisit.Frequency,
	}
}

func (visit *Visit) ToFullJson() VisitJson {
	visitJson := visit.ToJson()

	visitItemsJson := make([]VisitItemJson, len(visit.VisitItems))
	for i := range visit.VisitItems { 
		visitItemsJson[i] = visit.VisitItems[i].ToJson()
	}

	visitJson.VisitItems = visitItemsJson
	return visitJson
}
