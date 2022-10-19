package router

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/prybintsev/stakefish/internal/api"
	"github.com/prybintsev/stakefish/internal/config"
)

func StartHttpServer(ctx context.Context, logEntry *logrus.Entry, cfg config.AppConfig) error {
	router := gin.Default()

	aboutHandler := api.NewAboutHandler(logEntry, cfg)
	router.GET("/", aboutHandler.About)

	v1Group := router.Group("v1")
	toolsGroup := v1Group.Group("tools")
	lookupHandler := api.NewIPLookupHandler(logEntry)
	toolsGroup.GET("lookup", lookupHandler.Lookup)
	toolsGroup.POST("validate", lookupHandler.Validate)

	addr := fmt.Sprintf(":%d", cfg.Port)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: router,
	}
	logEntry.Info(fmt.Sprintf("starting server on %s", addr))
	go func() {
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err := srv.Shutdown(shutdownCtx)
		if err != nil {
			logEntry.WithError(err).Error("Could not gracefully shut down http server")
		}
	}()

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}
