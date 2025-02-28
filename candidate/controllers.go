package candidate

import (
	"go-gin-api/models"
	"io"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	// "strconv"
)

// CandidateController is the controller for candidate
type CandidateController struct{}
type reqBody struct {
	EmailMessageID     string `json:"email_message_id"`
	EmailUID           string `json:"email_uid"`
	AttachmentFilename string `json:"attachment_filename"`
	PersonalDetails    struct {
		FullName         string `json:"full_name"`
		MobileNumber     string `json:"mobile_number"`
		Email            string `json:"email"`
		Gender           string `json:"gender"`
		DateOfBirth      string `json:"date_of_birth"`
		Nationality      string `json:"nationality"`
		MaritalStatus    string `json:"marital_status"`
		CurrentLocation  string `json:"current_location"`
		PermanentAddress string `json:"permanent_address"`
	} `json:"personal_details"`
	ProfessionalExperience struct {
		TotalExperience struct {
			Years  int `json:"years"`
			Months int `json:"months"`
		} `json:"total_experience"`
		CurrentJob struct {
			Title    string `json:"title"`
			Employer struct {
				Name        string `json:"name"`
				Industry    string `json:"industry"`
				CompanySize string `json:"company_size"`
			} `json:"employer"`
			StartDate            string `json:"start_date"`
			CurrentlyWorkingHere bool   `json:"currently_working_here"`
			CurrentSalary        int    `json:"current_salary"`
			NoticePeriod         string `json:"notice_period"`
			EmploymentType       string `json:"employment_type"`
			JobResponsibilities  string `json:"job_responsibilities"`
			KeyAchievements      string `json:"key_achievements"`
		} `json:"current_job"`
		PreviousJobs []struct {
			Title            string `json:"title"`
			Company          string `json:"company"`
			StartDate        string `json:"start_date"`
			EndDate          string `json:"end_date"`
			ReasonForLeaving string `json:"reason_for_leaving"`
		} `json:"previous_jobs"`
	} `json:"professional_experience"`
	SkillsAndQualifications struct {
		KeySkills []struct {
			Skill string `json:"skill"`
		} `json:"key_skills"`
		TechnicalSkills []struct {
			Skill string `json:"skill"`
		} `json:"technical_skills"`
		LanguagesKnown []struct {
			Language         string `json:"language"`
			ProficiencyLevel string `json:"proficiency_level"`
		} `json:"languages_known"`
		Certifications []struct {
			Name                string `json:"name"`
			IssuingOrganization string `json:"issuing_organization"`
			DateOfIssuance      string `json:"date_of_issuance"`
			ExpirationDate      string `json:"expiration_date"`
		} `json:"certifications"`
	} `json:"skills_and_qualifications"`
	Education struct {
		HighestDegree struct {
			Degree              string `json:"degree"`
			FieldOfStudy        string `json:"field_of_study"`
			InstitutionName     string `json:"institution_name"`
			UniversityBoardName string `json:"university_board_name"`
			StartDate           string `json:"start_date"`
			EndDate             string `json:"end_date"`
			GPAPercentage       string `json:"gpa_percentage"`
		} `json:"highest_degree"`
	} `json:"education"`
	Preferences struct {
		PreferredJobTitle     string `json:"preferred_job_title"`
		PreferredLocation     string `json:"preferred_location"`
		ExpectedSalary        int    `json:"expected_salary"`
		Availability          string `json:"availability"`
		WillingnessToRelocate string `json:"willingness_to_relocate"`
	} `json:"preferences"`
	AdditionalDetails struct {
		LinkedInProfileURL          string `json:"linkedin_profile_url"`
		PersonalWebsitePortfolioURL string `json:"personal_website_portfolio_url"`
		ProfessionalSummary         string `json:"professional_summary"`
		AdditionalInformation       string `json:"additional_information"`
		AwardsAndAchievements       []struct {
			Name        string `json:"name"`
			Date        string `json:"date"`
			Description string `json:"description"`
		} `json:"awards_and_achievements"`
	} `json:"additional_details"`
}

// NewCandidateController returns a new candidate controller
func NewCandidateController() *CandidateController {
	return &CandidateController{}
}
func pointerToString(s string) *string {
	ps := new(string)
	*ps = s
	return ps
}

func pointerToBool(b bool) *bool {
	pb := new(bool)
	*pb = b
	return pb
}

// pointerToInt takes an int and returns a pointer to it. It is used for passing
// ints to functions that take a pointer to an int as an argument.
func pointerToInt(i int) *int {
	pi := new(int)
	*pi = i
	return pi
}

// reqBodyToCandidate takes a reqBody and converts it to a models.Candidate, while also checking if certain
// fields already exist in the database and using their IDs instead of creating new ones.
func reqBodyToCandidate(reqBody reqBody) models.Candidate {
	candidate := models.Candidate{
		EmailMessageID:     reqBody.EmailMessageID,
		EmailUID:           reqBody.EmailUID,
		AttachmentFilename: reqBody.AttachmentFilename,
		FullName:           reqBody.PersonalDetails.FullName,
		MobileNumber:       pointerToString(reqBody.PersonalDetails.MobileNumber),
		Email:              pointerToString(reqBody.PersonalDetails.Email),
		Gender:             pointerToString(reqBody.PersonalDetails.Gender),
		DateOfBirth:        pointerToString(reqBody.PersonalDetails.DateOfBirth),
		Nationality:        pointerToString(reqBody.PersonalDetails.Nationality),
		MaritalStatus:      pointerToString(reqBody.PersonalDetails.MaritalStatus),
		CurrentLocation:    pointerToString(reqBody.PersonalDetails.CurrentLocation),
		PermanentAddress:   pointerToString(reqBody.PersonalDetails.PermanentAddress),
		ProfessionalExperience: models.ProfessionalExperience{
			StartDate:            pointerToString(reqBody.ProfessionalExperience.CurrentJob.StartDate),
			CurrentlyWorkingHere: pointerToBool(reqBody.ProfessionalExperience.CurrentJob.CurrentlyWorkingHere),
			CurrentSalary:        pointerToInt(reqBody.ProfessionalExperience.CurrentJob.CurrentSalary),
			NoticePeriod:         pointerToString(reqBody.ProfessionalExperience.CurrentJob.NoticePeriod),
			EmploymentType:       pointerToString(reqBody.ProfessionalExperience.CurrentJob.EmploymentType),
			JobResponsibilities:  pointerToString(reqBody.ProfessionalExperience.CurrentJob.JobResponsibilities),
			KeyAchievements:      pointerToString(reqBody.ProfessionalExperience.CurrentJob.KeyAchievements),
			CandidateID:          0,
		},
		PreviousJobs: []models.PreviousJob{},

		Skills: []models.Skill{},

		Languages: []models.Language{},

		Certifications: []models.Certification{},

		Degree:                      pointerToString(reqBody.Education.HighestDegree.Degree),
		FieldOfStudy:                pointerToString(reqBody.Education.HighestDegree.FieldOfStudy),
		InstitutionName:             pointerToString(reqBody.Education.HighestDegree.InstitutionName),
		UniversityBoardName:         pointerToString(reqBody.Education.HighestDegree.UniversityBoardName),
		StartDate:                   pointerToString(reqBody.Education.HighestDegree.StartDate),
		EndDate:                     pointerToString(reqBody.Education.HighestDegree.EndDate),
		GPAPercentage:               pointerToString(reqBody.Education.HighestDegree.GPAPercentage),
		PreferredJobTitle:           pointerToString(reqBody.Preferences.PreferredJobTitle),
		PreferredLocation:           pointerToString(reqBody.Preferences.PreferredLocation),
		ExpectedSalary:              pointerToInt(reqBody.Preferences.ExpectedSalary),
		Availability:                pointerToString(reqBody.Preferences.Availability),
		WillingnessToRelocate:       pointerToString(reqBody.Preferences.WillingnessToRelocate),
		LinkedInProfileURL:          pointerToString(reqBody.AdditionalDetails.LinkedInProfileURL),
		PersonalWebsitePortfolioURL: pointerToString(reqBody.AdditionalDetails.PersonalWebsitePortfolioURL),
		ProfessionalSummary:         pointerToString(reqBody.AdditionalDetails.ProfessionalSummary),
		AdditionalInformation:       pointerToString(reqBody.AdditionalDetails.AdditionalInformation),
		AwardAchievements:           []models.AwardAchievement{},
	}

	for _, a := range reqBody.AdditionalDetails.AwardsAndAchievements {
		candidate.AwardAchievements = append(candidate.AwardAchievements, models.AwardAchievement{
			Name:        a.Name,
			Date:        pointerToString(a.Date),
			Description: pointerToString(a.Description),
		})
	}

	for _, lang := range reqBody.SkillsAndQualifications.LanguagesKnown {
		existingLang, err := CheckIfLanguageExists(lang.Language, lang.ProficiencyLevel)
		if err != nil {
			existingLang = &models.Language{}
		}
		candidate.Languages = append(candidate.Languages, models.Language{
			Language:         lang.Language,
			ProficiencyLevel: pointerToString(lang.ProficiencyLevel),
		})

		candidate.Languages[len(candidate.Languages)-1].ID = existingLang.ID

	}

	for _, skill := range reqBody.SkillsAndQualifications.KeySkills {
		skill.Skill = strings.ToLower(skill.Skill)
		id, err := CheckIfSkillExists(skill.Skill)
		if err != nil {
			id = &models.Skill{}
		}
		candidate.Skills = append(candidate.Skills, models.Skill{
			Skill: skill.Skill,
		})
		candidate.Skills[len(candidate.Skills)-1].ID = id.ID
	}
	for _, skill := range reqBody.SkillsAndQualifications.TechnicalSkills {
		skill.Skill = strings.ToLower(skill.Skill)
		id, err := CheckIfSkillExists(skill.Skill)
		if err != nil {
			id = &models.Skill{}
		}
		candidate.Skills = append(candidate.Skills, models.Skill{

			Skill: skill.Skill,
		})
		candidate.Skills[len(candidate.Skills)-1].ID = id.ID
	}
	for _, pj := range reqBody.ProfessionalExperience.PreviousJobs {
		candidate.PreviousJobs = append(candidate.PreviousJobs, models.PreviousJob{
			Title:            pointerToString(pj.Title),
			Company:          pointerToString(pj.Company),
			StartDate:        pointerToString(pj.StartDate),
			EndDate:          pointerToString(pj.EndDate),
			ReasonForLeaving: pointerToString(pj.ReasonForLeaving),
			CandidateID:      0,
		})
	}
	for _, cert := range reqBody.SkillsAndQualifications.Certifications {
		candidate.Certifications = append(candidate.Certifications, models.Certification{
			Name:                cert.Name,
			IssuingOrganization: pointerToString(cert.IssuingOrganization),
			DateOfIssuance:      pointerToString(cert.DateOfIssuance),
			ExpirationDate:      pointerToString(cert.ExpirationDate),
		})
	}
	return candidate
}

// CreateCandidate creates a new candidate
func (controller *CandidateController) CreateCandidate(c *gin.Context) {

	var reqBody reqBody

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	candidate := reqBodyToCandidate(reqBody)
	CreateCandidate(candidate)
	c.JSON(http.StatusCreated, candidate)
}

// ReadLogFile returns the contents of the log file
func (controller *CandidateController) ReadLogFile(c *gin.Context) {
	file, err := os.Open("/home/ubuntu/application.log")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	var logText []byte
	logText, err = io.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.String(http.StatusOK, string(logText))
}

// FindAllCandidates returns all candidates
func (controller *CandidateController) FindAllCandidates(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	offset := (page - 1) * pageSize

	candidates := FindAllCandidates(offset, pageSize)
	c.JSON(http.StatusOK, candidates)
}

// GetNumberOfPages returns the number of pages of candidates available
func (controller *CandidateController) GetNumberOfPages(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	numberOfRecords := CountCandidates()
	numberOfPages := int(math.Ceil(float64(numberOfRecords) / float64(pageSize)))
	c.JSON(http.StatusOK, gin.H{"numberOfPages": numberOfPages})
}

// FindCandidateByID returns a candidate by id
func (controller *CandidateController) FindCandidateByID(c *gin.Context) {
	id := c.Param("id")
	candidate, err := FindCandidateByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, candidate)
}

// GET /candidate/byPhone
func (controller *CandidateController) FindCandidateByPhone(c *gin.Context) {
	phone := c.Query("phone")
	candidate, err := FindCandidateByPhone(phone)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, candidate)
}

// UpdateCandidate updates a candidate
func (controller *CandidateController) UpdateCandidate(c *gin.Context) {
	id := c.Param("id")
	var reqBody reqBody

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	candidate := reqBodyToCandidate(reqBody)
	candidateID, _ := strconv.ParseUint(id, 10, 64)
	candidate.ID = uint(candidateID)
	UpdateCandidate(candidate)
	c.JSON(http.StatusOK, candidate)
}

// DeleteCandidate deletes a candidate
func (controller *CandidateController) DeleteCandidate(c *gin.Context) {
	id := c.Param("id")
	candidate, err := FindCandidateByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	DeleteCandidate(candidate)
	c.JSON(http.StatusNoContent, gin.H{})
}

// GetAllEmailUIDs returns a list of all the emailUIDs
func (controller *CandidateController) GetAllEmailUIDs(c *gin.Context) {
	emailUIDs, err := GetAllEmailUIDs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, emailUIDs)
}

// AddEmailUID adds an email UID if it doesn't exist already
func (controller *CandidateController) AddEmailUID(c *gin.Context) {
	var reqBody struct {
		EmailUID string `json:"emailUID"`
	}
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	emailUID := reqBody.EmailUID
	err := AddEmailUIDIfNotExists(emailUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{})
}
