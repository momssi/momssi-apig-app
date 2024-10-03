package app

import (
	"context"
	"log"
	"momssi-apig-app/config"
	"momssi-apig-app/internal/database"
	"momssi-apig-app/internal/logger"
	"momssi-apig-app/internal/server"
	"sync"
)

type App struct {
	cfg *config.EnvConfig
	srv *server.Gin
	db  *database.MySqlClient
}

func NewApplication(ctx context.Context) *App {

	cfg, err := config.LoadEnvConfig()
	if err != nil {
		log.Fatalf("fail to read config, err : %v", err)
	}

	if err := logger.SlogInit(cfg.Logger); err != nil {
		log.Fatalf("fail to init slog err : %v", err)
	}

	db, err := database.NewMysqlClient(cfg.Mysql)
	if err != nil {
		log.Fatalf("fail to connect mysql client, err : %v", err)
	}

	srv := server.NewGinServer(cfg.Server)

	app := &App{
		cfg: cfg,
		srv: srv,
		db:  db,
	}

	return app
}

func (a *App) Start(wg *sync.WaitGroup) {
	a.srv.Run(wg)
}

func (a *App) Stop(ctx context.Context) {
	a.srv.Shutdown(ctx)
}
