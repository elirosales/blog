package database

import (
	"fmt"
	"log"

	"github.com/elizabethrosales/blog/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// Migrate executes db migrations
func Migrate(c config.Config) error {
	dbURL := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", c.Database.Username, c.Database.Password, c.Database.Host, c.Database.Port, c.Database.DBName)
	m, err := migrate.New("file://migrations", dbURL)

	if err != nil {
		return err
	}

	if err = m.Up(); err != nil {
		log.Printf("UP_ERROR: %v", err.Error())
		if err.Error() != "no change" {
			return err
		}
	}

	return nil
}
