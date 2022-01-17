package applicant

import (
	"go-gin-api/models"
	"go-gin-api/utils/bpm"
	"go-gin-api/utils/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

type applicant struct {
	BusinessKey string `json:"businessKey"`
}

type attachApplicantsRequest struct {
	ParentID string `json:"parentId"`
	Applicants []applicant `json:"applicants"`
}


func AttachApplicants(c *gin.Context) {
	var applicants attachApplicantsRequest
	err := c.ShouldBindJSON(&applicants)

	if err != nil {
		c.Error(common.AppError{Message:"Invalid Request",Code:http.StatusBadRequest})
		return
	}
	ProcessExists,err := bpm.CheckProcessExists(applicants.ParentID)

	if err != nil {
		c.Error(common.AppError{Message:"Some internal server error occured",Code:http.StatusInternalServerError})
		return
	}

	if !ProcessExists{
		c.Error(common.AppError{Message:"Parent process does not exist",Code:http.StatusFailedDependency})
		return
	}

	for _,applicant := range applicants.Applicants {
		data,err := bpm.CreateProcessInstance("applicant",applicant.BusinessKey)
		if err != nil {
			c.Error(common.AppError{Message:"Failed to start Bpm process in camunda",Code:http.StatusInternalServerError})
			return 
		}

		pI := models.ProcessInstance{
			BusinessKey: data.BusinessKey,
			ProcessID:   data.ID,
			ProcessType: "applicant",
			Owner:       applicants.ParentID,
			OwnerType:   "Process",
		}
		bpm.InsertProcessInstance(pI)
	}
	c.JSON(http.StatusCreated,gin.H{"message":"Created"})
}
