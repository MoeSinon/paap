package service

import (
	"log"

	"github.com/MoeSinon/paap/db"
	"github.com/MoeSinon/paap/service/user"
)

type Mapper struct {
	UserSvc user.Service
	Logger  *log.Logger
}

// New initializes all services and helpers.
func New(client *db.Client, logger *log.Logger) *Mapper {
	return &Mapper{
		UserSvc: user.New(client),
		Logger:  logger,
	}
}
