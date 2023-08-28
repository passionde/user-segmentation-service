package app

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/passionde/user-segmentation-service/config"
	v1 "github.com/passionde/user-segmentation-service/internal/controller/http/v1"
	"github.com/passionde/user-segmentation-service/internal/repo"
	"github.com/passionde/user-segmentation-service/internal/service"
	"github.com/passionde/user-segmentation-service/pkg/csvwriter"
	"github.com/passionde/user-segmentation-service/pkg/httpserver"
	"github.com/passionde/user-segmentation-service/pkg/postgres"
	"github.com/passionde/user-segmentation-service/pkg/secure"
	"github.com/passionde/user-segmentation-service/pkg/validator"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func Run(configPath string) {
	// Config
	cfg, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Logger
	SetLogrus(cfg.Log.Level)

	// Database
	log.Info("Initializing postgres...")
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.MaxPoolSize))
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	// Repositories
	log.Info("Initializing repositories...")
	repositories := repo.NewRepositories(pg)

	// Services
	log.Info("Initializing services...")
	deps := service.ServicesDependencies{
		Repos:     repositories,
		APISecure: secure.NewSecure(cfg.Secure.Salt),
		CSVWrite:  csvwriter.NewCsvWriter("reports"),
	}
	services := service.NewServices(deps)

	// Echo
	log.Info("Initializing handlers and routes...")
	handler := echo.New()
	// setup handler validator as lib validator
	handler.Validator = validator.NewCustomValidator()
	v1.NewRouter(handler, services)

	// HTTP server
	log.Info("Starting http server...")
	log.Debugf("Server port: %s", cfg.HTTP.Port)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	log.Info("Configuring graceful shutdown...")
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		log.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Graceful shutdown
	log.Info("Shutting down...")
	err = httpServer.Shutdown()
	if err != nil {
		log.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
