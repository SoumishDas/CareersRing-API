package routes

import (
	"go-gin-api/authentication"
	"go-gin-api/candidate"
	"go-gin-api/hcm"
	"go-gin-api/masterData"

	"github.com/gin-contrib/cors"
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

	// CORS configuration
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{
		"GET",
		"POST",
		"PUT",
		"DELETE",
		"OPTIONS",
	}
	config.AllowHeaders = []string{
		"Origin",
		"Content-Type",
		"Authorization",
	}
	// Apply the CORS middleware
	Router.Use(cors.New(config))

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
	api.GET("/candidate/byPhone", candidateCtrl.FindCandidateByPhone)

	api.GET("/candidate/:id", candidateCtrl.FindCandidateByID)
	api.PUT("/candidate/:id", candidateCtrl.UpdateCandidate)
	api.GET("/candidate/numPages", candidateCtrl.GetNumberOfPages)
	api.DELETE("/candidate/:id", candidateCtrl.DeleteCandidate)
	api.GET("/log", candidateCtrl.ReadLogFile)
	api.POST("/user", authentication.CreateUserController)
	api.GET("/emailuids", candidateCtrl.GetAllEmailUIDs)
	api.POST("/emailuid", candidateCtrl.AddEmailUID)

	masterDataCtrl := masterData.NewMasterDataController()
	api.POST("/masterData/candidates", masterDataCtrl.CreateMasterCandidate)
	api.GET("/masterData/candidates", masterDataCtrl.FindAllMasterCandidates)
	api.GET("/masterData/candidates/numPages", masterDataCtrl.GetNumberOfPages)
	api.GET("/masterData/candidates/:id", masterDataCtrl.FindMasterCandidateByID)
	api.PUT("/masterData/candidates/:id", masterDataCtrl.UpdateMasterCandidate)
	api.DELETE("/masterData/candidates/:id", masterDataCtrl.DeleteMasterCandidate)
	api.GET("/masterData/bridge/:candidateID", masterDataCtrl.BridgeOldCandidate)

	api.GET("/masterData/skills", masterDataCtrl.GetAllSkills)
	api.GET("/masterData/languages", masterDataCtrl.GetAllLanguages)
	api.GET("/masterData/jobTitles", masterDataCtrl.GetAllJobTitles)
	api.GET("/masterData/locations", masterDataCtrl.GetAllLocations)
	api.GET("/masterData/industries", masterDataCtrl.GetAllIndustries)

	return Router
}
