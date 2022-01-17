package requirement

import (
	"go-gin-api/db"
	"go-gin-api/models"
)

func createDbRequirement(r models.Requirement) models.Requirement {
	db.DB.Create(&r)
	return r
}