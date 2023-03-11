package database

import (
	"log"
	"os"
	"rosenzu/database/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB
var err error

func SetupDb() {
	dsn := os.Getenv("DATABASE_URL")

	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("FAILED TO CONNECT TO DB")
	} else {
		log.Println("CONNECTED TO DB")
	}
	Db.AutoMigrate(&model.Line{}, &model.Element{}, &model.Relation{}, &model.Coordinate{}, &model.OperationalPoint{})
	DbSeed()
}
