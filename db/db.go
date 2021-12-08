package db

import (
	"fmt"
	"net/url"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB gorm.DB

// Connecting to db
func ConnectDB() {
	dsn := url.URL{
	User: url.UserPassword("fastapi_user", "apppassword"),
	Host:     fmt.Sprintf("%s:%d","18.140.204.101" , 5432),
	Path:     "fastapi",
	Scheme:   "postgres",
	RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}
	db, err := gorm.Open("postgres", dsn.String())
	if err!= nil{
		println("Error Connecting to db")
	}
	DB = *db
	
}

func CloseDB(){
	err := DB.Close()
	if err!= nil{
		println("Failed to Close DB")
	}
}