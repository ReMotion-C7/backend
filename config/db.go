package config

import (
	"ReMotion-C7/constant"
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
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false,
	})

	if err != nil {
		log.Fatalln(constant.ErrConnectingDatabase)
	}

	database = db

}

func CheckConnection() {

	sqlDb, err := database.DB()
	if err != nil {
		log.Fatalf(string(constant.ErrGetSQLInstance))
	}

	err = sqlDb.Ping()
	if err != nil {
		log.Fatalln(constant.ErrPingDatabase)
	}

	log.Println(constant.SuccessConnectDatabase)

}

func GetDatabase() *gorm.DB {

	return database

}
