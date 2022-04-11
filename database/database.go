package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DatabaseInstance struct {
	DB *gorm.DB
}

var Database DatabaseInstance

func ConnectDatabase() {
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), "3306", os.Getenv("DB_NAME"))
	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database \n", err.Error())
		os.Exit(2)
	}

	log.Println("Connected to the database successfully")
	db.Logger = logger.Default.LogMode(logger.Info)

	Database = DatabaseInstance{DB: db}
}
