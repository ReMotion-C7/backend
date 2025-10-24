package config

import (
	"ReMotion-C7/constant"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

func ConnectDatabase() {

	dbUrl := LoadEnvConfig("DIRECT_URL")

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dbUrl,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	
	if err != nil {
		log.Fatalf(constant.ErrConnectingDatabase)
	}

	database = db

}

func CheckConnection() {

	sqlDb, err := database.DB()
	if err != nil {
		log.Fatalf(constant.ErrGetSQLInstance)
	}

	err = sqlDb.Ping()
	if err != nil {
		log.Fatalf(constant.ErrPingDatabase)
	}

	fmt.Println(constant.SuccessConnectDatabase)

}

func GetDatabase() *gorm.DB {

	return database

}
