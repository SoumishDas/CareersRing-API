package masterData

import (
	"go-gin-api/candidate"
	"go-gin-api/models"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// MasterDataController is the controller for MasterCandidate
type MasterDataController struct{}

// NewMasterDataController returns a new controller instance
func NewMasterDataController() *MasterDataController {
	return &MasterDataController{}
}

// CreateMasterCandidate handles POST /masterData/candidates
func (ctrl *MasterDataController) CreateMasterCandidate(c *gin.Context) {
	var mc models.MasterCandidate
	if err := c.ShouldBindJSON(&mc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	created := CreateMasterCandidate(mc)
	c.JSON(http.StatusCreated, created)
}

// FindAllMasterCandidates handles GET /masterData/candidates
func (ctrl *MasterDataController) FindAllMasterCandidates(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	offset := (page - 1) * pageSize

	candidates := FindAllMasterCandidates(offset, pageSize)
	c.JSON(http.StatusOK, candidates)
}

// GetNumberOfPages handles GET /masterData/candidates/numberOfPages
func (ctrl *MasterDataController) GetNumberOfPages(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	total := CountMasterCandidates()
	numPages := int(math.Ceil(float64(total) / float64(pageSize)))
	c.JSON(http.StatusOK, gin.H{"numberOfPages": numPages})
}

// FindMasterCandidateByID handles GET /masterData/candidates/:id
func (ctrl *MasterDataController) FindMasterCandidateByID(c *gin.Context) {
	idStr := c.Param("id")
	idVal, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	mc, findErr := FindMasterCandidateByID(uint(idVal))
	if findErr != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": findErr.Error()})
		return
	}
	c.JSON(http.StatusOK, mc)
}

// UpdateMasterCandidate handles PUT /masterData/candidates/:id
func (ctrl *MasterDataController) UpdateMasterCandidate(c *gin.Context) {
	idStr := c.Param("id")
	idVal, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var mc models.MasterCandidate
	if bindErr := c.ShouldBindJSON(&mc); bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": bindErr.Error()})
		return
	}

	mc.ID = uint(idVal)
	updated := UpdateMasterCandidate(mc)
	c.JSON(http.StatusOK, updated)
}

// DeleteMasterCandidate handles DELETE /masterData/candidates/:id
func (ctrl *MasterDataController) DeleteMasterCandidate(c *gin.Context) {
	idStr := c.Param("id")
	idVal, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	mc, findErr := FindMasterCandidateByID(uint(idVal))
	if findErr != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": findErr.Error()})
		return
	}

	DeleteMasterCandidate(mc)
	c.JSON(http.StatusNoContent, gin.H{})
}

// BridgeOldCandidate fetches an old candidate, converts to MasterCandidate,
// then maps MasterCandidate -> final JSON that matches your form fields.
func (ctrl *MasterDataController) BridgeOldCandidate(c *gin.Context) {
	oldID := c.Param("candidateID")

	// 1. Fetch from old parser DB
	oldCandidate, err := candidate.FindCandidateByID(oldID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Old candidate not found"})
		return
	}

	// 2. Convert oldCandidate -> MasterCandidate
	mc := convertOldCandidateToMaster(oldCandidate)

	// 3. Convert MasterCandidate -> form-friendly JSON
	formData := convertMasterCandidateToFormData(mc)

	// 4. Return JSON
	c.JSON(http.StatusOK, formData)
}

// --- Dictionary Endpoints ---

// GetAllSkills handles GET /masterData/skills
func (ctrl *MasterDataController) GetAllSkills(c *gin.Context) {
	skills := GetAllMasterSkills()
	c.JSON(http.StatusOK, skills)
}

// GetAllLanguages handles GET /masterData/languages
func (ctrl *MasterDataController) GetAllLanguages(c *gin.Context) {
	languages := GetAllMasterLanguages()
	c.JSON(http.StatusOK, languages)
}

// GetAllJobTitles handles GET /masterData/jobTitles
func (ctrl *MasterDataController) GetAllJobTitles(c *gin.Context) {
	titles := GetAllMasterJobTitles()
	c.JSON(http.StatusOK, titles)
}

// GetAllLocations handles GET /masterData/locations
func (ctrl *MasterDataController) GetAllLocations(c *gin.Context) {
	locations := GetAllMasterLocations()
	c.JSON(http.StatusOK, locations)
}

// GetAllIndustries handles GET /masterData/industries
func (ctrl *MasterDataController) GetAllIndustries(c *gin.Context) {
	industries := GetAllMasterIndustries()
	c.JSON(http.StatusOK, industries)
}
