package authentication

import (
	"go-gin-api/utils/bpm"
)


type LoginService interface {
	LoginUser(email string, password string) (bool,uint64)
}
type loginInformation struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewLoginService() LoginService{
	return &loginInformation{}
}

func (info *loginInformation) LoginUser(email string, password string) (bool,uint64) {
	
	user,err := bpm.VerifyUser(loginInformation{Username: email,Password:password })

	if err!=nil{
		panic(err)
	}
	if user.Authenticated{
		return true,1
	}else{
		return false,0
	}

	
}