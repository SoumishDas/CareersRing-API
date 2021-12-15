package main

import (
	"go-gin-api/db"
	router "go-gin-api/router"
	"log"
	"runtime"

	"go-gin-api/models"

	"github.com/gin-gonic/gin"
)

var (Router *gin.Engine)

func main() {
	runtime.GOMAXPROCS(2)
	db.ConnectDB()
	Router = router.GetRouter()
	models.MigrateDB(&db.DB)
	
	Router.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
		"message": 2,
		})
	})

	log.Fatal(Router.Run(":5000"))
}
