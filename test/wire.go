//go:build wireinject

package test

import (
	"context"
	"log"

	"github.com/MoeSinon/paap/config"
	"github.com/MoeSinon/paap/db"
	"github.com/MoeSinon/paap/logger"
	"github.com/google/wire"
)

type App struct {
	Cfg    *config.Config
	Client *db.Client
	Ctx    context.Context
	Logger *log.Logger
}

func New(cfg *config.Config, client *db.Client, logger *log.Logger) *App {
	return &App{cfg, client, context.Background(), logger}
}

func Initialize() (*App, error) {
	wire.Build(config.New, db.New, logger.New, New)
	return &App{}, nil
}

func (a *App) Close() error {
	return a.Client.Close()
}
