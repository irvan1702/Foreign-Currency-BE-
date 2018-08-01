package db

import (
	"currency-exchange/models"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

var db *gorm.DB
var err error

func Init() {
	user := "postgres"
	password := ""
	host := "localhost"
	port := "5432"
	database := "ForeignExchange"

	dbinfo := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		user,
		password,
		host,
		port,
		database,
	)

	db, err = gorm.Open("postgres", dbinfo)
	if err != nil {
		log.Println("Failed to connect to server")
		panic(err)
	}

	log.Println("Database Connected")

	if !db.HasTable(&models.ForexHistory{}) {
		err := db.Create(&models.ForexHistory{})
		if err != nil {
			log.Println("Table already exists")
		}
	}

	db.AutoMigrate(&models.ForexHistory{})
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	db.Close()
}
