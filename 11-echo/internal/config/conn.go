package config

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (cfg *Config) openConn() (*gorm.DB, error) {
	dsn := fmt.Sprintf("postgres://%v:%v@%v:%v/%v", cfg.DB_User, cfg.DB_Password, cfg.DB_Url, cfg.DB_Port, cfg.DB_Name)
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}

	sqldb, err := db.DB()
	if err != nil {
		return nil, err
	}
	if err := sqldb.Ping(); err != nil {
		return nil, err
	}

	sqldb.SetMaxOpenConns(cfg.DbMaxOpenConn)
	sqldb.SetMaxIdleConns(cfg.DbMaxIdleConn)

	idletime, err := time.ParseDuration(cfg.DbMaxIdleTime)
	if err != nil {
		return nil, fmt.Errorf("failed parsing maxidletime: %v", err)
	}
	sqldb.SetConnMaxIdleTime(idletime)

	return db, nil
}
