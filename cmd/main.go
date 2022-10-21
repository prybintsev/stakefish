package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"

	_ "github.com/prybintsev/stakefish/docs"
	"github.com/prybintsev/stakefish/internal/config"
	"github.com/prybintsev/stakefish/internal/db/migrations"
	"github.com/prybintsev/stakefish/internal/db/postgres"
	"github.com/prybintsev/stakefish/internal/router"
	"github.com/prybintsev/stakefish/internal/version"
)

const x = "bla"

// @title  Stakefish API

// @host      localhost:3000
// @BasePath  /

func main() {
	logEntry := logrus.WithField("version", version.Version)
	logEntry.Info("Starting stakefish application")

	cfg := config.Init(logEntry)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go listenToSignals(cancel)

	db := postgres.ConnectToPostgres(cfg)
	err := migrations.MigrateUp(db, cfg, logEntry)
	if err != nil {
		logrus.WithError(err).Error("Failed to perform DB migrations")
		return
	}

	err = router.StartHttpServer(ctx, logEntry, cfg, db)
	if err != nil {
		logrus.WithError(err).Error("HTTP server has stopped unexpectedly")
		return
	}
	logrus.Info("Exiting")
}

func listenToSignals(cancel context.CancelFunc) {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	<-shutdown
	logrus.Info("Gracefully shutting down the http server")
	cancel()
}
