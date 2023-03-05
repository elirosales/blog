package main

import (
	"github.com/elizabethrosales/blog/config"
	"github.com/elizabethrosales/blog/database"
	"github.com/elizabethrosales/blog/router"
	"github.com/elizabethrosales/blog/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	// initialize dependencies
	godotenv.Load()
	c := config.New()

	err := database.Migrate(*c)
	if err != nil {
		panic(err.Error())
	}

	db, err := database.Initialize(*c)
	if err != nil {
		panic(err.Error())
	}

	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			panic(err)
		}

		sqlDB.Close()
	}()

	log := logrus.New()
	svc := service.New(log, db)

	r := router.New(svc, log)
	r.Run(c.API.Port)
}
