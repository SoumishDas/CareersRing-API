package masterData

import (
	"go-gin-api/db"
	"go-gin-api/models"
	"strconv"
)

// CreateMasterCandidate inserts a new MasterCandidate into DB
func CreateMasterCandidate(mc models.MasterCandidate) models.MasterCandidate {
	db.DB.Create(&mc)
	return mc
}

// FindAllMasterCandidates returns a paginated list
func FindAllMasterCandidates(offset, pageSize int) []models.MasterCandidate {
	var candidates []models.MasterCandidate
	db.DB.Preload("ProfessionalExperience").
		Preload("PreviousJobs").
		Preload("Skills").
		Preload("Languages").
		Preload("Certifications").
		Preload("AwardAchievements").
		Preload("JobTitles").
		Preload("PreferredLocations").
		Limit(pageSize).Offset(offset).Find(&candidates)
	return candidates
}

// CountMasterCandidates returns total number of MasterCandidate records
func CountMasterCandidates() int {
	var count int
	db.DB.Model(&models.MasterCandidate{}).Count(&count)
	return count
}

// FindMasterCandidateByID fetches a MasterCandidate by ID
func FindMasterCandidateByID(id uint) (models.MasterCandidate, error) {
	var mc models.MasterCandidate
	result := db.DB.Preload("ProfessionalExperience").
		Preload("PreviousJobs").
		Preload("Skills").
		Preload("Languages").
		Preload("Certifications").
		Preload("AwardAchievements").
		Preload("JobTitles").
		Preload("PreferredLocations").
		First(&mc, id)

	return mc, result.Error
}

// UpdateMasterCandidate updates an existing MasterCandidate
func UpdateMasterCandidate(mc models.MasterCandidate) models.MasterCandidate {
	db.DB.Save(&mc)
	return mc
}

// DeleteMasterCandidate removes a MasterCandidate
func DeleteMasterCandidate(mc models.MasterCandidate) {
	db.DB.Delete(&mc)
}

// convertOldCandidateToMaster transforms the old parser-based candidate
// into the new MasterCandidate structure. Adjust field mappings as needed.
func convertOldCandidateToMaster(oldC models.Candidate) models.MasterCandidate {
	// Basic fields
	newMC := models.MasterCandidate{
		FullName:         oldC.FullName,
		Email:            oldC.Email,
		MobileNumber:     oldC.MobileNumber,
		Gender:           oldC.Gender,
		DateOfBirth:      oldC.DateOfBirth,
		Nationality:      oldC.Nationality,
		MaritalStatus:    oldC.MaritalStatus,
		CurrentLocation:  oldC.CurrentLocation,
		PermanentAddress: oldC.PermanentAddress,

		// Example: If you have new toggles like AutoUpdateLinkedIn or EducationType,
		// you can set them to false / nil or guess a default if not in oldC.
		AutoUpdateLinkedIn: nil,
		EducationType:      nil,

		// Education fields
		Degree:              oldC.Degree,
		FieldOfStudy:        oldC.FieldOfStudy,
		InstitutionName:     oldC.InstitutionName,
		UniversityBoardName: oldC.UniversityBoardName,
		StartDate:           oldC.StartDate,
		EndDate:             oldC.EndDate,
		GPAPercentage:       oldC.GPAPercentage,

		// Preferences
		ExpectedSalary:        oldC.ExpectedSalary,
		Availability:          oldC.Availability,
		WillingnessToRelocate: oldC.WillingnessToRelocate,

		// Additional
		LinkedInProfileURL:          oldC.LinkedInProfileURL,
		PersonalWebsitePortfolioURL: oldC.PersonalWebsitePortfolioURL,
		ProfessionalSummary:         oldC.ProfessionalSummary,
		AdditionalInformation:       oldC.AdditionalInformation,
	}

	// 1-to-1 ProfessionalExperience
	// (assuming oldC.ProfessionalExperience has a similar shape)
	newMC.ProfessionalExperience = models.MasterProfessionalExperience{
		StartDate:            oldC.ProfessionalExperience.StartDate,
		CurrentlyWorkingHere: oldC.ProfessionalExperience.CurrentlyWorkingHere,
		CurrentSalary:        oldC.ProfessionalExperience.CurrentSalary,
		NoticePeriod:         oldC.ProfessionalExperience.NoticePeriod,
		EmploymentType:       oldC.ProfessionalExperience.EmploymentType,
		JobResponsibilities:  oldC.ProfessionalExperience.JobResponsibilities,
		KeyAchievements:      oldC.ProfessionalExperience.KeyAchievements,
	}

	// 1-to-many PreviousJobs
	newMC.PreviousJobs = make([]models.MasterPreviousJob, len(oldC.PreviousJobs))
	for i, pj := range oldC.PreviousJobs {
		newMC.PreviousJobs[i] = models.MasterPreviousJob{
			Title:            pj.Title,
			Company:          pj.Company,
			StartDate:        pj.StartDate,
			EndDate:          pj.EndDate,
			ReasonForLeaving: pj.ReasonForLeaving,
		}
	}

	// many-to-many Skills -> MasterSkill
	// oldC.Skills is []Skill (where skill.Skill is a string?), new = []MasterSkill
	newMC.Skills = make([]models.MasterSkill, len(oldC.Skills))
	for i, sk := range oldC.Skills {
		newMC.Skills[i] = models.MasterSkill{
			Name: sk.Skill, // If old is "Skill", new is "Name"
		}
	}

	// many-to-many Languages -> MasterLanguage
	newMC.Languages = make([]models.MasterLanguage, len(oldC.Languages))
	for i, lg := range oldC.Languages {
		newMC.Languages[i] = models.MasterLanguage{
			Name:             lg.Language,
			ProficiencyLevel: lg.ProficiencyLevel,
		}
	}

	// many-to-many Certifications
	newMC.Certifications = make([]models.MasterCertification, len(oldC.Certifications))
	for i, cert := range oldC.Certifications {
		newMC.Certifications[i] = models.MasterCertification{
			Name:                cert.Name,
			IssuingOrganization: cert.IssuingOrganization,
			DateOfIssuance:      cert.DateOfIssuance,
			ExpirationDate:      cert.ExpirationDate,
		}
	}

	// many-to-many AwardAchievements
	newMC.AwardAchievements = make([]models.MasterAwardAchievement, len(oldC.AwardAchievements))
	for i, award := range oldC.AwardAchievements {
		newMC.AwardAchievements[i] = models.MasterAwardAchievement{
			Name:        award.Name,
			Date:        award.Date,
			Description: award.Description,
		}
	}

	// If oldC has a single PreferredJobTitle, you might turn that into a slice of MasterJobTitle
	if oldC.PreferredJobTitle != nil && *oldC.PreferredJobTitle != "" {
		newMC.JobTitles = []models.MasterJobTitle{
			{Title: *oldC.PreferredJobTitle},
		}
	}

	// If oldC has a single PreferredLocation, same approach:
	if oldC.PreferredLocation != nil && *oldC.PreferredLocation != "" {
		newMC.PreferredLocations = []models.MasterLocation{
			{Name: *oldC.PreferredLocation},
		}
	}

	return newMC
}
func convertMasterCandidateToFormData(mc models.MasterCandidate) map[string]interface{} {
	// Helper to safely read pointers
	safeString := func(p *string) string {
		if p == nil {
			return ""
		}
		return *p
	}
	safeBool := func(p *bool) bool {
		if p == nil {
			return false
		}
		return *p
	}
	// safeInt := func(p *int) int {
	// 	if p == nil {
	// 		return 0
	// 	}
	// 	return *p
	// }

	// phone, email, name, etc.
	phone := safeString(mc.MobileNumber)
	email := safeString(mc.Email)
	name := mc.FullName // if FullName is not a pointer, just use it directly

	// Skills: convert []MasterSkill to []string
	var skills []string
	for _, sk := range mc.Skills {
		skills = append(skills, sk.Name)
	}

	// Industries: if you stored them in mc.Industries, similarly convert
	// or if you used "PreferredJobTitle" as a single string, you might skip

	// Example: Convert ProfessionalExperience fields
	annualSalary := ""
	if mc.ProfessionalExperience.CurrentSalary != nil {
		annualSalary = strconv.Itoa(*mc.ProfessionalExperience.CurrentSalary)
	}

	noticePeriod := safeString(mc.ProfessionalExperience.NoticePeriod)

	// If you store join date as year/month in the MasterCandidate, map them
	joinYear := ""
	joinMonth := ""
	// e.g., if you stored them as separate pointers in ProfessionalExperience or somewhere else
	// or skip if not needed

	// If you have a boolean for "currentlyWorking"
	currentlyWorking := safeBool(mc.ProfessionalExperience.CurrentlyWorkingHere)

	// Convert languages, jobTitles, preferredLocations, etc. to slices of strings
	var languages []string
	for _, lg := range mc.Languages {
		languages = append(languages, lg.Name)
	}

	var jobTitles []string
	for _, jt := range mc.JobTitles {
		jobTitles = append(jobTitles, jt.Title)
	}

	var preferredCities []string
	for _, loc := range mc.PreferredLocations {
		preferredCities = append(preferredCities, loc.Name)
	}

	// Example: for "industries" if you used MasterCandidate.Skills for them, or a separate field
	var industries []string
	for _, ind := range mc.Industries {
		industries = append(industries, ind.Name)
	}
	// (If you have a separate array in the struct. Adjust as needed.)

	// "gender", "highestQualification", "university", etc.
	gender := safeString(mc.Gender)
	highestQualification := safeString(mc.Degree)    // or if you used "highestQualification"
	university := safeString(mc.UniversityBoardName) // or if you store it differently

	// Education type
	educationType := safeString(mc.EducationType)

	// LinkedIn
	linkedinUrl := safeString(mc.LinkedInProfileURL)
	autoUpdateLinkedIn := safeBool(mc.AutoUpdateLinkedIn)

	// totalExperienceYears / totalExperienceMonths if you stored them
	// or skip if you store them in a single field

	// Build final map
	return map[string]interface{}{
		"phone":                phone,
		"email":                email,
		"name":                 name,
		"skills":               skills,
		"annualSalary":         annualSalary,
		"noticePeriod":         noticePeriod,
		"joinYear":             joinYear,
		"joinMonth":            joinMonth,
		"currentlyWorking":     currentlyWorking,
		"languages":            languages,
		"jobTitles":            jobTitles,
		"preferredCities":      preferredCities,
		"industries":           industries,
		"gender":               gender,
		"highestQualification": highestQualification,
		"university":           university,
		"educationType":        educationType,
		"linkedinUrl":          linkedinUrl,
		"autoUpdateLinkedIn":   autoUpdateLinkedIn,
		// etc. for all fields your form expects
	}
}

// --- Dictionary Helpers ---

// GetAllMasterSkills returns all skills from the dictionary table
func GetAllMasterSkills() []models.MasterSkill {
	var skills []models.MasterSkill
	db.DB.Find(&skills)
	return skills
}

// GetAllMasterLanguages returns all languages
func GetAllMasterLanguages() []models.MasterLanguage {
	var languages []models.MasterLanguage
	db.DB.Find(&languages)
	return languages
}

// GetAllMasterJobTitles returns all job titles
func GetAllMasterJobTitles() []models.MasterJobTitle {
	var titles []models.MasterJobTitle
	db.DB.Find(&titles)
	return titles
}

// GetAllMasterLocations returns all preferred locations
func GetAllMasterLocations() []models.MasterLocation {
	var locations []models.MasterLocation
	db.DB.Find(&locations)
	return locations
}

// GetAllMasterIndustries returns all industries
func GetAllMasterIndustries() []models.MasterIndustry {
	var industries []models.MasterIndustry
	db.DB.Find(&industries)
	return industries
}
