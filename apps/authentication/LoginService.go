package authentication

import (
	"errors"
	"go-gin-api/db"
	"go-gin-api/models"

	"github.com/jinzhu/gorm"
)


type LoginService interface {
	LoginUser(email string, password string) (bool,uint64)
}
type loginInformation struct {

}

func NewLoginService() LoginService{
	return &loginInformation{}
}

func (info *loginInformation) LoginUser(email string, password string) (bool,uint64) {
	user := models.User{Email:email,Password:password}
	if result := db.DB.Where(&user).First(&models.User{}); result.Error != nil {
		// error handling...
		if errors.Is(result.Error, gorm.ErrRecordNotFound){
			return false,999
		}
		return false,9999
	  }
	return true,user.ID
}