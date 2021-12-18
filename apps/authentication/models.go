package authentication

import (
	"go-gin-api/db"
	"go-gin-api/models"
)




func CreateUser(user models.User) models.User {
	db.DB.Create(&user)
	return user
}


func UpdateUser(user models.User) models.User{
	db.DB.Save(user)
	return user
}


func DeleteUser(user models.User) {
	db.DB.Delete(user)
}

func FindAllUser() []models.User {
	var users []models.User
	db.DB.Find(&users)
	return users
}