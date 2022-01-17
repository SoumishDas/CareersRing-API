package requirement

import (
	"go-gin-api/models"
	"go-gin-api/utils/bpm"
	"go-gin-api/utils/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createRequirementRequest struct {
	BusinessKey string `json:"businessKey"`
}

func CreateRequirement(c *gin.Context) {
	var body createRequirementRequest
	err := c.ShouldBindJSON(&body)

	if err != nil{
		c.Error(common.AppError{Message: "Invalid Request",Code:http.StatusBadRequest})
		return
	}
	data,err := bpm.CreateProcessInstance("test",body.BusinessKey)

	
	if err != nil {
		c.Error(err)
		return
	}
	pI := models.ProcessInstance{BusinessKey: data.BusinessKey,ProcessID: data.ID,ProcessType: "test",Owner:"soumish",OwnerType: "User"}
	pI = bpm.InsertProcessInstance(pI)

	c.JSON(http.StatusOK,pI)
}


