package db

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/lib/pq" // Postgres driver
)

var DB gorm.DB

// Connecting to db
func ConnectDB(dbType string) {
	//var err error
	if dbType == "postgres" {
		// db, err := gorm.Open("postgres", "user=postgres dbname=test sslmode=disable password=123456")
		dbName := "Test"
		if os.Getenv("ENV") == "Production" {
			dbName = "Prod"
		}
		db, err := gorm.Open("postgres", "user=postgres host=43.205.211.80 dbname="+dbName+" sslmode=disable password=chikoo123")
		if err != nil {
			log.Fatal("Error Connecting to db")
		}
		DB = *db
	} else if dbType == "sqlite3" {
		db, err := gorm.Open("sqlite3", "./test.db")
		if err != nil {
			log.Fatal("Error Connecting to db")
		}
		DB = *db
	} else {
		log.Fatal("Invalid db type")
	}

}

// CloseDB closes the database connection. It should be called when the program
// is exiting to ensure that the connection is properly closed.
func CloseDB() {
	err := DB.Close()
	if err != nil {
		println("Failed to Close DB")
	}
}
