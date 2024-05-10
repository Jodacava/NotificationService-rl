package postgres

import (
	"NotificationService-rl/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

func NewPostgres() (*gorm.DB, error) {
	DB, err := gorm.Open("postgres", config.DbStringConnection)
	if err != nil {
		return nil, err
	}
	DB.DB().SetMaxIdleConns(2)
	DB.DB().SetMaxOpenConns(10)
	DB.DB().SetConnMaxLifetime(time.Second * 60)
	DB.LogMode(true)

	return DB, nil
}
