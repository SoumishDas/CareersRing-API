package main

import (
	"go-gin-api/db"
	router "go-gin-api/router"
	"log"

	"go-gin-api/models"

	"github.com/gin-gonic/gin"
)

var (Router *gin.Engine)

func main() {
	Router = router.GetRouter()
	


	db.ConnectDB()
	
	models.MigrateDB(&db.DB)
	
	

	log.Fatal(Router.Run(":5000"))
}
