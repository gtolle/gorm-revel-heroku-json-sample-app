package controllers

import "github.com/revel/revel"
import "github.com/gtolle/gorm-revel-heroku-json-sample-app/app"
import "github.com/gtolle/gorm-revel-heroku-json-sample-app/app/models"

type VisitsController struct {
	*revel.Controller
}

func (c VisitsController) Index() revel.Result {
	var visits []models.Visit
	app.Gorm.Preload("User").Preload("Stylist.User").Preload("RecurringVisit").Order("date asc").Where("visit_state = ?","open").Find(&visits)

	visitsJson := make([]models.VisitJson, len(visits))
	for i := range visits { 
		visitsJson[i] = visits[i].ToJson() 
	}

	return c.RenderJson(visitsJson)
}

func (c VisitsController) Show(id string) revel.Result {
	var visit models.Visit
	app.Gorm.Preload("User").Preload("Stylist.User").Preload("RecurringVisit").Preload("VisitItems").First(&visit, id)
	return c.RenderJson(visit.ToFullJson())
}
