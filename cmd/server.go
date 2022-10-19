package main

import (
	"context"
	"github.com/prybintsev/stakefish/internal/db/migrations"
	"github.com/prybintsev/stakefish/internal/db/postgres"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"

	"github.com/prybintsev/stakefish/internal/config"
	"github.com/prybintsev/stakefish/internal/router"
	"github.com/prybintsev/stakefish/internal/version"
)

func listenToSignals(cancel context.CancelFunc) {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	<-shutdown
	logrus.Info("Gracefully shutting down the http server")
	cancel()
}

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
