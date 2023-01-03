package database

import (
	"fmt"
	"github.com/oktapascal/app-barayya/bootstraps"
	"github.com/oktapascal/app-barayya/models/domain"
	"github.com/oktapascal/app-barayya/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

func NewMysql(config bootstraps.Config) *gorm.DB {
	var (
		host     = config.Get("DB_HOST")
		port     = config.Get("DB_PORT")
		dbname   = config.Get("DB_NAME")
		user     = config.Get("DB_USER")
		password = config.Get("DB_PASSWORD")
		mode     = config.Get("MODE")
	)

	logType := logger.Silent
	if mode == "DEVELOPMENT" {
		logType = logger.Info
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			Colorful:                  false,
			IgnoreRecordNotFoundError: true,
			LogLevel:                  logType,
		},
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})

	utils.PanicIfError(err)

	mysqlDb, err := db.DB()

	utils.PanicIfError(err)

	mysqlDb.SetMaxOpenConns(70)
	mysqlDb.SetMaxIdleConns(70)
	mysqlDb.SetConnMaxIdleTime(10 * time.Minute)
	mysqlDb.SetConnMaxLifetime(15 * time.Minute)

	err = db.AutoMigrate(&domain.User{}, &domain.Karyawan{}, &domain.Session{})

	utils.PanicIfError(err)

	return db
}
