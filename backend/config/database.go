package config

import (
	"log"
	"time"

	mysqlDriver "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadDB() {

	log.Printf("log url %s\n", ENV.DBUrl)

	cfg := mysqlDriver.Config{
		User:                 ENV.DBUsername,
		Passwd:               ENV.DBPassword,
		Net:                  "tcp",
		Addr:                 ENV.DBUrl,
		DBName:               ENV.DBDatabase,
		ParseTime:            true,
		Loc:                  time.UTC,
		AllowNativePasswords: true,
	}

	dsn := cfg.FormatDSN()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed connect database:", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed get sqlDB:", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	DB = db
}
