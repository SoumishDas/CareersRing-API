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

