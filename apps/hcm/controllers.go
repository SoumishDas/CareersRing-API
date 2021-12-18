package hcm

import (
	"go-gin-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


func GetItems(c *gin.Context){
	results := FindAll()
	c.JSON(http.StatusOK,results)
}

func CreateItem(c *gin.Context){
	requestBody := models.Item{}
	c.Bind(&requestBody)
	
	item := Create(requestBody)
	print(requestBody.Name)
	c.JSON(http.StatusCreated,item)
}
func DeleteItem(c *gin.Context){
	requestBody := models.Item{}
	c.Bind(&requestBody)
	
	Delete(requestBody)
	print(requestBody.Name)
	c.JSON(http.StatusNoContent,gin.H{})
}


type BpmKey struct {
	key string `json:"key"`
}
func CreateBpm(c *gin.Context){
	reqBody := BpmKey{}
	err := c.ShouldBindJSON(&reqBody)
	if err!= nil {
		c.JSON(http.StatusBadRequest,gin.H{"msg":"No Key specified","error":err.Error()})
	}
}
