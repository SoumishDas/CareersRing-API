package routes

import (
	"go-gin-api/authentication"
	"go-gin-api/candidate"
	"go-gin-api/hcm"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

func sleep(c *gin.Context) {
	for i := 0; i < 10000000000; {
		i++
	}
	c.JSON(200, gin.H{"msg": "success"})
}
func GetRouter() *gin.Engine {

	Router := gin.Default()
	Router.Use(gindump.Dump())
	Router.POST("/login", authentication.MainLoginHandler)
	Router.POST("/refresh", authentication.MainRefreshHandler)

	api := Router.Group("/")
	//api.Use(middleware.AuthorizeJWT())
	api.GET("/Item", hcm.GetItems)
	api.GET("/sleep", sleep)

	api.POST("/Item", hcm.CreateItem)
	api.DELETE("/Item", hcm.DeleteItem)

	// Candidate methods
	// Candidate methods
	candidateCtrl := candidate.NewCandidateController()
	api.POST("/candidate", candidateCtrl.CreateCandidate)
	api.GET("/candidates", candidateCtrl.FindAllCandidates)
	api.GET("/candidate/:id", candidateCtrl.FindCandidateByID)
	// api.PUT("/candidate/:id", candidateCtrl.UpdateCandidate)
	api.DELETE("/candidate/:id", candidateCtrl.DeleteCandidate)
	api.GET("/log", candidateCtrl.ReadLogFile)
	api.POST("/user", authentication.CreateUserController)

	return Router
}
