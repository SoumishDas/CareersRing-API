package routes

import (
	"go-gin-api/apps/authentication"
	"go-gin-api/apps/hcm"
	"go-gin-api/utils/bpm"

	"github.com/gin-gonic/gin"
)

func sleep (c *gin.Context) {
	for i:=0 ;i<10000000000;{
		i++
	}
	c.JSON(200,gin.H{"msg":"success"})
}
func GetRouter() *gin.Engine {
	
	Router := gin.Default()
	//Router.Use(gindump.Dump())
	Router.POST("/login",authentication.MainLoginHandler)
	Router.POST("/refresh",authentication.MainRefreshHandler)
	
	api := Router.Group("/")
	//api.Use(middleware.AuthorizeJWT())
	api.GET("/Item", hcm.GetItems)
	api.GET("/sleep",sleep)
	api.GET("/start",func(c *gin.Context){
		err := bpm.CreateProcessInstance("test")
		print(err)
		if !err{
			println("Success")
			c.JSON(200,gin.H{"success":"true"})
		}
	})
	api.POST("/Item", hcm.CreateItem)
	api.DELETE("/Item", hcm.DeleteItem)

	// User methods
	
	api.POST("/user",authentication.CreateUserController)


	return Router
}
