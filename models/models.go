package models

import "github.com/jinzhu/gorm"

func MigrateDB(db *gorm.DB){
	db.AutoMigrate(&Item{},&User{})
}