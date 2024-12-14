package main

import (
	"go-gin-api/db"
	router "go-gin-api/router"
	"log"
	"os"
	"runtime"

	"go-gin-api/models"

	"github.com/gin-gonic/gin"
)

var (
	Router *gin.Engine
)

func main() {
	runtime.GOMAXPROCS(2)
	if os.Getenv("ENV") == "Production" {
		gin.SetMode(gin.ReleaseMode)
	}

	//test
	db.ConnectDB("postgres")
	Router = router.GetRouter()
	models.MigrateDB(&db.DB)

	log.Fatal(Router.Run(":5000"))

}
