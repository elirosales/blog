package service

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Service struct {
	log *logrus.Logger
	db  *gorm.DB
}

// New returns a new service
func New(l *logrus.Logger, db *gorm.DB) *Service {
	return &Service{
		log: l,
		db:  db,
	}
}
