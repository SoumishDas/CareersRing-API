package bpm

import (
	"go-gin-api/db"
	"go-gin-api/models"
)

func InsertProcessInstance(pI models.ProcessInstance) models.ProcessInstance {
	db.DB.Create(&pI)
	return pI
}

func CheckProcessExists(ID string) (bool,error){
	print(ID)
	var value models.ProcessInstance
	r := db.DB.Where("process_id = ?", ID).First(&value) 
		 
         
         
	exists := r.RowsAffected > 0
	
	if r.Error != nil {
			return false, r.Error
		}
	return exists,nil
	
}