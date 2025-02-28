package models

import "github.com/jinzhu/gorm"

type ProfessionalExperience struct {
	gorm.Model
	StartDate            *string
	CurrentlyWorkingHere *bool
	CurrentSalary        *int
	NoticePeriod         *string
	EmploymentType       *string
	JobResponsibilities  *string
	KeyAchievements      *string
	CandidateID          uint64 `gorm:"not null"`
}

type PreviousJob struct {
	gorm.Model
	Title            *string
	Company          *string
	StartDate        *string
	EndDate          *string
	ReasonForLeaving *string
	CandidateID      uint64 `gorm:"not null"`
}

type Skill struct {
	gorm.Model
	Skill string `gorm:"not null"`
}

type Language struct {
	gorm.Model
	Language         string `gorm:"not null"`
	ProficiencyLevel *string
}

type Certification struct {
	gorm.Model
	Name                string `gorm:"not null"`
	IssuingOrganization *string
	DateOfIssuance      *string
	ExpirationDate      *string
}

type AdditionalDetails struct {
	gorm.Model
	LinkedInProfileURL          *string
	PersonalWebsitePortfolioURL *string
	ProfessionalSummary         *string
}

type AwardAchievement struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Date        *string
	Description *string
}
type Candidate struct {
	gorm.Model
	EmailUID                    string
	EmailMessageID              string
	AttachmentFilename          string
	FullName                    string `gorm:"not null"`
	MobileNumber                *string
	Email                       *string
	Gender                      *string
	DateOfBirth                 *string
	Nationality                 *string
	MaritalStatus               *string
	CurrentLocation             *string
	PermanentAddress            *string
	ProfessionalExperience      ProfessionalExperience `gorm:"ForeignKey:CandidateID"`
	PreviousJobs                []PreviousJob          `gorm:"ForeignKey:CandidateID"`
	Skills                      []Skill                `gorm:"many2many:candidate_skills;"`
	Languages                   []Language             `gorm:"many2many:candidate_languages;"`
	Certifications              []Certification        `gorm:"many2many:candidate_certifications;"`
	Degree                      *string
	FieldOfStudy                *string
	InstitutionName             *string
	UniversityBoardName         *string
	StartDate                   *string
	EndDate                     *string
	GPAPercentage               *string
	PreferredJobTitle           *string
	PreferredLocation           *string
	ExpectedSalary              *int
	Availability                *string
	WillingnessToRelocate       *string
	LinkedInProfileURL          *string
	PersonalWebsitePortfolioURL *string
	ProfessionalSummary         *string
	AdditionalInformation       *string
	AwardAchievements           []AwardAchievement `gorm:"many2many:candidate_award_achievements;"`
}

type EmailUID struct {
	gorm.Model
	UID string `gorm:"not null;unique"`
}
