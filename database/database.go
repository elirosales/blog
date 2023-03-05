package database

import (
	"fmt"
	"time"

	"github.com/elizabethrosales/blog/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Initialize initializes database connection
func Initialize(c config.Config) (*gorm.DB, error) {
	var dsn string
	if c.Database.DSN != "" {
		dsn = c.Database.DSN
	} else {
		dsn = fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable connect_timeout=5", c.Database.Host, c.Database.Username, c.Database.Password, c.Database.DBName, c.Database.Port)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetConnMaxLifetime(5 * time.Minute)
	return db, err
}
