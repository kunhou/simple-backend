package data

import (
	"log"
	"time"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(debug bool, dataConf *DatabaseConf) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dataConf.GetConnection()), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	if debug {
		db = db.Debug()
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	if dataConf.MaxConn > 0 {
		sqlDB.SetMaxOpenConns(dataConf.MaxConn)
	}
	if dataConf.MaxIdle > 0 {
		sqlDB.SetMaxIdleConns(dataConf.MaxIdle)
	}
	if dataConf.LifeTime > 0 {
		sqlDB.SetConnMaxLifetime(time.Duration(dataConf.LifeTime) * time.Second)
	}

	return db
}
