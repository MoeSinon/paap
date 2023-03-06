//go:build wireinject

package app

import (
	"github.com/MoeSinon/paap/config"
	"github.com/MoeSinon/paap/db"
	"github.com/MoeSinon/paap/logger"
	"github.com/MoeSinon/paap/server"
	"github.com/MoeSinon/paap/service"
	"github.com/google/wire"
)

func Initialize() (*App, error) {
	wire.Build(config.New, db.New, service.New, logger.New, server.New, New)
	return &App{}, nil
}
