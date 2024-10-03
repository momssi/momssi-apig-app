package app

import (
	"context"
	"log"
	"momssi-apig-app/config"
	"momssi-apig-app/internal/logger"
	"momssi-apig-app/internal/server"
	"sync"
)

type App struct {
	cfg *config.EnvConfig
	srv *server.Gin
}

func NewApplication(ctx context.Context) *App {

	cfg, err := config.LoadEnvConfig()
	if err != nil {
		log.Fatalf("fail to read config, err : %v", err)
	}

	if err := logger.SlogInit(cfg.Logger); err != nil {
		log.Fatalf("fail to init slog err : %v", err)
	}

	srv := server.NewGinServer(cfg.Server)

	app := &App{
		cfg: cfg,
		srv: srv,
	}

	return app
}

func (a *App) Start(wg *sync.WaitGroup) {
	a.srv.Run(wg)
}

func (a *App) Stop(ctx context.Context) {
	a.srv.Shutdown(ctx)
}
