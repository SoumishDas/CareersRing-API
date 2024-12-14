package candidate

import (
	"go-gin-api/db"
	"go-gin-api/models"

	// "github.com/jinzhu/gorm"
	// "gorm.io/gorm/clause"
	"log"
)

// CreateCandidate creates a new candidate
func CreateCandidate(candidate models.Candidate) models.Candidate {
	log.Println("Creating candidate...")
	db.DB.Debug().Create(&candidate)
	log.Println("Candidate created:", candidate.ID)
	return candidate
}

// FindAllCandidates returns all candidates
func FindAllCandidates(offset int, pageSize int) []models.Candidate {
	var candidates []models.Candidate
	db.DB.Preload("ProfessionalExperience").Preload("PreviousJobs").Preload("Skills").Preload("Languages").Preload("Certifications").Preload("AwardAchievements").
		Limit(pageSize).Offset(offset).Find(&candidates)
	return candidates
}

// FindCandidateByID returns a candidate by id
func FindCandidateByID(id string) (models.Candidate, error) {
	var candidate models.Candidate
	result := db.DB.Where("id = ?", id).Preload("ProfessionalExperience").Preload("PreviousJobs").Preload("Skills").Preload("Languages").Preload("Certifications").Preload("AwardAchievements").First(&candidate)
	if result.Error != nil {
		return candidate, result.Error
	}
	return candidate, nil
}

// UpdateCandidate updates a candidate
func UpdateCandidate(candidate models.Candidate) models.Candidate {
	db.DB.Save(&candidate)
	return candidate
}

// DeleteCandidate deletes a candidate
func DeleteCandidate(candidate models.Candidate) {
	db.DB.Delete(candidate)
}

// CheckIfSkillExists checks if a particular skill exists
func CheckIfSkillExists(skillName string) (*models.Skill, error) {
	var skill models.Skill
	result := db.DB.Where("skill = ?", skillName).First(&skill)
	if result.Error != nil {
		return nil, result.Error
	}
	return &skill, nil
}

// CheckIfLanguageExists checks if a particular language exists
func CheckIfLanguageExists(languageName, proficiency string) (*models.Language, error) {
	var language models.Language
	result := db.DB.Where("language = ? AND proficiency_level = ?", languageName, proficiency).First(&language)
	if result.Error != nil {
		return nil, result.Error
	}
	return &language, nil
}
