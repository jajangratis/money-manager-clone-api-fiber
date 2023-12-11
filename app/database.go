package app

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

func OpenConnection() *gorm.DB {
	godotenv.Load()
	dsn := "host=" + os.Getenv("DB_HOST") + " user=" + os.Getenv("DB_USERNAME") + " password=" + os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_SELECT") + " port=" + os.Getenv("DB_PORT") + " sslmode=disable TimeZone=Asia/Bangkok"
	loggerInfo := logger.Error
	if os.Getenv("ENVIROMENT") != "local" {
		loggerInfo = logger.Error
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:      logger.Default.LogMode(loggerInfo),
		PrepareStmt: true,
	})
	if err != nil {
		panic(err)
	}
	return db
}
