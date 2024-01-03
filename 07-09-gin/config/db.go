package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenConn(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatal("error open postgres:", err)
	}

	sqldb, err := db.DB()
	if err != nil {
		log.Fatal("error get sqldb: ", err)
	}
	if err := sqldb.Ping(); err != nil {
		log.Fatal("error ping connection: ", err)
	}

	// // Database Pooling
	// DB.DB().SetMaxIdleConns(20)
	// DB.DB().SetMaxOpenConns(200)
	// DB.DB().SetConnMaxLifetime(45 * time.Second)
	return db
}
