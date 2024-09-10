package setting

import (
	"github.com/sirupsen/logrus"
	"github.com/yehey-1030/household-account-book/go/constants/config"
	"gorm.io/driver/mysql"
	gormlib "gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func SetUpDB() *gormlib.DB {
	var db *gormlib.DB
	var err error

	db, err = gormlib.Open(mysql.Open(config.DBConfigInfo.DataSource), &gormlib.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		logrus.Panicf("fail to open MYSQL DB Config [%s]", err.Error())
	}

	dbConfig, err := db.DB()
	if err != nil {
		logrus.Panicf("fail to open MYSQL DB Config [%s]", err.Error())
	}
	dbConfig.SetMaxOpenConns(10)
	dbConfig.SetMaxIdleConns(5)
	dbConfig.SetConnMaxIdleTime(4 * time.Hour)

	return db
}
