package router

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/prybintsev/stakefish/internal/db/lookup"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/prybintsev/stakefish/internal/api"
	"github.com/prybintsev/stakefish/internal/config"
)

func StartHttpServer(ctx context.Context, logEntry *logrus.Entry, cfg config.AppConfig, db *sql.DB) error {
	router := gin.Default()

	aboutHandler := api.NewAboutHandler(logEntry, cfg)
	router.GET("/", aboutHandler.AppInfo)

	lookupRepo := lookup.NewLookupRepo(db)
	lookupHandler := api.NewIPLookupHandler(logEntry, lookupRepo)
	v1Group := router.Group("v1")
	v1Group.GET("history", lookupHandler.History)

	toolsGroup := v1Group.Group("tools")
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
