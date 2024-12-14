package models

import "github.com/jinzhu/gorm"

// MigrateDB creates the tables in the database.
func MigrateDB(db *gorm.DB){
	db.AutoMigrate(
		
		&ProfessionalExperience{},
		&PreviousJob{},
		&Skill{},
		&Language{},
		&Certification{},
		&AwardAchievement{},
		&AdditionalDetails{},
		&Candidate{},

	)
}