package models

import "github.com/jinzhu/gorm"

// MasterCandidate holds all the "master" fields
type MasterCandidate struct {
	gorm.Model

	// Personal / Contact Info
	FullName         string `gorm:"not null"`
	Email            *string
	MobileNumber     *string
	Gender           *string
	DateOfBirth      *string
	Nationality      *string
	MaritalStatus    *string
	CurrentLocation  *string
	PermanentAddress *string

	// Additional toggles / new fields
	AutoUpdateLinkedIn *bool
	EducationType      *string

	// One-to-one
	ProfessionalExperience MasterProfessionalExperience `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	// One-to-many
	PreviousJobs []MasterPreviousJob `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	// Many-to-many (dictionary-based)
	Skills             []MasterSkill            `gorm:"many2many:master_candidate_skills;"`
	Languages          []MasterLanguage         `gorm:"many2many:master_candidate_languages;"`
	Certifications     []MasterCertification    `gorm:"many2many:master_candidate_certifications;"`
	AwardAchievements  []MasterAwardAchievement `gorm:"many2many:master_candidate_award_achievements;"`
	JobTitles          []MasterJobTitle         `gorm:"many2many:master_candidate_job_titles;"`
	PreferredLocations []MasterLocation         `gorm:"many2many:master_candidate_locations;"`
	Industries         []MasterIndustry         `gorm:"many2many:master_candidate_industries;"` // <--- NEW

	// Education
	Degree              *string
	FieldOfStudy        *string
	InstitutionName     *string
	UniversityBoardName *string
	StartDate           *string
	EndDate             *string
	GPAPercentage       *string

	// Preferences
	ExpectedSalary        *int
	Availability          *string
	WillingnessToRelocate *string

	// Additional
	LinkedInProfileURL          *string
	PersonalWebsitePortfolioURL *string
	ProfessionalSummary         *string
	AdditionalInformation       *string
}

// MasterProfessionalExperience (one-to-one with MasterCandidate)
type MasterProfessionalExperience struct {
	gorm.Model

	StartDate            *string
	CurrentlyWorkingHere *bool
	CurrentSalary        *int
	NoticePeriod         *string
	EmploymentType       *string
	JobResponsibilities  *string
	KeyAchievements      *string

	MasterCandidateID uint
}

// MasterPreviousJob (one-to-many)
type MasterPreviousJob struct {
	gorm.Model

	Title            *string
	Company          *string
	StartDate        *string
	EndDate          *string
	ReasonForLeaving *string

	MasterCandidateID uint
}

// Dictionary tables for many-to-many usage

type MasterSkill struct {
	gorm.Model
	Name string `gorm:"uniqueIndex;not null"`
}

type MasterLanguage struct {
	gorm.Model
	Name             string `gorm:"uniqueIndex;not null"`
	ProficiencyLevel *string
}

type MasterCertification struct {
	gorm.Model
	Name                string `gorm:"not null"`
	IssuingOrganization *string
	DateOfIssuance      *string
	ExpirationDate      *string
}

type MasterAwardAchievement struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Date        *string
	Description *string
}

type MasterJobTitle struct {
	gorm.Model
	Title string `gorm:"uniqueIndex;not null"`
}

type MasterLocation struct {
	gorm.Model
	Name string `gorm:"uniqueIndex;not null"`
}

// MasterIndustry is the NEW dictionary table for storing industries
type MasterIndustry struct {
	gorm.Model
	Name string `gorm:"uniqueIndex;not null"`
}
