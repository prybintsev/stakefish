package main

import (
	"context"
	"github.com/prybintsev/stakefish/internal/version"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"

	"github.com/prybintsev/stakefish/internal/router"
)

func listenToSignals(cancel context.CancelFunc) {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	<-shutdown
	log.Info("Gracefully shutting down the http server")
	cancel()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go listenToSignals(cancel)
	log.WithField("version", version.Version).Info("Starting stakefish application")
	err := router.StartHttpServer(ctx)
	if err != nil {
		log.WithError(err).Fatal("Authentication server has stopped unexpectedly")
	}
	log.Info("Exiting")
}
