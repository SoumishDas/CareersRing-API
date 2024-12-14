package hcm

import (
	"go-gin-api/db"
	"go-gin-api/models"
)








func Create(item models.Item) models.Item {
	db.DB.Create(&item)
	return item
}


func Update(item models.Item) models.Item{
	db.DB.Save(item)
	return item
}


func Delete(item models.Item) {
	db.DB.Delete(item)
}

func FindAll() []models.Item {
	var items []models.Item
	db.DB.Find(&items)
	return items
}


