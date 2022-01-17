package routes

import (
	"go-gin-api/apps/applicant"
	"go-gin-api/apps/authentication"
	"go-gin-api/apps/hcm"
	"go-gin-api/apps/requirement"
	"go-gin-api/middleware"

	"github.com/gin-contrib/cors"
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
	corsConfig := cors.DefaultConfig()

	corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	// To be able to send tokens to the server.
	corsConfig.AllowCredentials = true

	// OPTIONS method for ReactJS
	corsConfig.AddAllowMethods("OPTIONS")

	// Register the middleware
	Router.Use(cors.New(corsConfig))
	// Use Error Handler
	Router.Use(middleware.JSONAppErrorReporter())

	//Router.Use(gindump.Dump())
	Router.POST("/login",authentication.MainLoginHandler)
	Router.POST("/refresh",authentication.MainRefreshHandler)
	
	api := Router.Group("/")
	//api.Use(middleware.AuthorizeJWT())
	api.GET("/Item", hcm.GetItems)
	api.GET("/sleep",sleep)
	
	api.POST("/Item", hcm.CreateItem)
	api.DELETE("/Item", hcm.DeleteItem)

	// User methods
	
	api.POST("/user",authentication.CreateUserController)

	api.POST("/applicant/attach",applicant.AttachApplicants)

	api.POST("/requirement/create",requirement.CreateRequirement)


	return Router
}
