package authentication

import (
	"errors"
	"fmt"
	"go-gin-api/db"
	"go-gin-api/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//login contorller interface
type LoginController interface {
	Login(ctx *gin.Context) Tokens
}

type loginController struct {
	loginService LoginService
	jWtService   JWTService
}

func LoginHandler(loginService LoginService,
	jWtService JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jWtService:   jWtService,
	}
}

//Login credential
type LoginCredentials struct {
	Email    string `json:"username" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

//Refresh credential
type tokenReqBody struct {
	RefreshToken string `json:"refresh_token"`
}

func (controller *loginController) Login(ctx *gin.Context) Tokens {
	var credential LoginCredentials
	err := ctx.ShouldBindJSON(&credential)
	if err != nil {
		println("No data Found")
		ctx.JSON(http.StatusNotAcceptable,gin.H{"error_msg":"request was not proper","error":err.Error()})
		return Tokens{}
		
	}
	println("got credentials")
	println(credential.Email)
	
	isUserAuthenticated,userId := controller.loginService.LoginUser(credential.Email, credential.Password)
	if isUserAuthenticated {
		println(userId)
		return controller.jWtService.GenerateToken(credential.Email, true,userId)

	}
	ctx.JSON(http.StatusNotAcceptable,gin.H{"error_msg":"User not found"})
	return Tokens{}
}

func CreateUserController(c *gin.Context){
	var user models.User
	c.Bind(user)
	CreateUser(user)
}

func MainLoginHandler(ctx *gin.Context){
	
	var loginController LoginController = LoginHandler(NewLoginService(),JWTAuthService())
	token := loginController.Login(ctx)
	if token.AccessToken != "" {
		ctx.JSON(http.StatusOK, gin.H{
			"access_token": token.AccessToken,
			"refresh_token": token.RefreshToken,
		})
	}
}

func MainRefreshHandler(ctx *gin.Context){
	
	tokenReq := tokenReqBody{}
	ctx.BindJSON(&tokenReq)
	if tokenReq.RefreshToken != "" {
		token, err := JWTAuthService().ValidateToken(tokenReq.RefreshToken)
		if err!= nil {
			println(err)
			ctx.JSON(http.StatusUnauthorized,gin.H{"err_msg":err.Error()})
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Get the user record from database or
			// run through your business logic to verify if the user can log in
			uid,err := strconv.ParseInt(claims["sub"].(string), 10, 64) 
			if err!= nil {
				println("error")
			}
			user := models.User{ID: uint64(uid) }
			result := db.DB.Where(&user).First(&models.User{})
			if result.Error != nil {
				// error handling...
				if errors.Is(result.Error, gorm.ErrRecordNotFound){
					ctx.JSON(http.StatusOK, gin.H{"error_msg":"User not found"})
				}
			  }
			  
			token := JWTAuthService().GenerateToken(user.Email,true,user.ID)
			
			
			
			ctx.JSON(http.StatusOK, gin.H{
				"access_token": token.AccessToken,
				"refresh_token": token.RefreshToken,
			})
			} else {
			fmt.Println(err)
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
	} else {
		
		ctx.JSON(http.StatusUnauthorized, nil)
	}
}