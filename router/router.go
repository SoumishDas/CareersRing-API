package routes

import (
	"go-gin-api/authentication"
	"go-gin-api/hcm"
	"go-gin-api/middleware"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)


func GetRouter() *gin.Engine {
	Router := gin.Default()
	Router.Use(gindump.Dump())
	Router.POST("/login",authentication.MainLoginHandler)
	Router.POST("/refresh",authentication.MainRefreshHandler)
	
	api := Router.Group("/")
	api.Use(middleware.AuthorizeJWT())
	api.GET("/Item", hcm.GetItems)
	api.POST("/Item", hcm.CreateItem)
	api.DELETE("/Item", hcm.DeleteItem)

	// User methods
	
	api.POST("/user",authentication.CreateUserController)


	return Router
}
