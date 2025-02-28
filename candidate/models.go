package candidate

import (
	"go-gin-api/db"
	"go-gin-api/models"

	// "github.com/jinzhu/gorm"
	// "gorm.io/gorm/clause"
	"log"

	"github.com/jinzhu/gorm"
)

// CreateCandidate creates a new candidate
func CreateCandidate(candidate models.Candidate) models.Candidate {
	var emailUID models.EmailUID
	if err := db.DB.Where("uid = ?", candidate.EmailUID).First(&emailUID).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			db.DB.Create(&models.EmailUID{UID: candidate.EmailUID})
		} else {
			log.Println("Error checking email UID:", err)
		}
	}
	log.Println("Creating candidate...")
	db.DB.Create(&candidate)
	log.Println("Candidate created:", candidate.ID)
	return candidate
}

// CountCandidates returns the number of candidates in the database
func CountCandidates() int {
	var count int
	db.DB.Model(&models.Candidate{}).Count(&count)
	return count
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

func FindCandidateByPhone(phone string) (*models.Candidate, error) {
	var candidate models.Candidate
	err := db.DB.Where("mobile_number = ?", phone).Preload("ProfessionalExperience").
		Preload("PreviousJobs").Preload("Skills").Preload("Languages").
		Preload("Certifications").Preload("AwardAchievements").
		First(&candidate).Error
	if err != nil {
		return nil, err
	}
	return &candidate, nil
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

// CheckIfEmailUIDExists checks if the particular emailUID already exists or not
func CheckIfEmailUIDExists(emailUID string) (bool, error) {
	var count int
	result := db.DB.Model(&models.Candidate{}).Where("email_uid = ?", emailUID).Count(&count)
	if result.Error != nil {
		return false, result.Error
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}

// GetAllEmailUIDs returns a list of all the emailUIDs
func GetAllEmailUIDs() ([]string, error) {
	var emailUIDs []string
	result := db.DB.Model(&models.EmailUID{}).Pluck("uid", &emailUIDs)
	if result.Error != nil {
		return nil, result.Error
	}
	return emailUIDs, nil
}

// AddEmailUIDIfNotExists adds an email UID if it doesn't exist already
func AddEmailUIDIfNotExists(emailUID string) error {
	var existingEmailUID models.EmailUID
	if err := db.DB.Where("uid = ?", emailUID).First(&existingEmailUID).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			db.DB.Create(&models.EmailUID{UID: emailUID})
		} else {
			return err
		}
	}
	return nil
}
