package app

import (
	"context"
	"log"
	"momssi-apig-app/config"
	"momssi-apig-app/internal/logger"
	"sync"
)

type App struct {
	cfg *config.EnvConfig
}

func NewApplication(ctx context.Context) *App {

	cfg, err := config.LoadEnvConfig()
	if err != nil {
		log.Fatalf("fail to read config, err : %v", err)
	}

	if err := logger.SlogInit(cfg.Logger); err != nil {
		log.Fatalf("fail to init slog err : %v", err)
	}

	app := &App{
		cfg: cfg,
	}

	return app
}

func (a *App) Start(wg *sync.WaitGroup) {
}

func (a *App) Stop(ctx context.Context) {
}
