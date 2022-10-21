package api

import (
	"github.com/prybintsev/stakefish/internal/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/prybintsev/stakefish/internal/config"
	"github.com/prybintsev/stakefish/internal/version"
)

type AboutHandler struct {
	logEntry *logrus.Entry
	cfg      config.AppConfig
}

func NewAboutHandler(logEntry *logrus.Entry, cfg config.AppConfig) AboutHandler {
	return AboutHandler{logEntry: logEntry, cfg: cfg}
}

// AppInfo godoc
// @Summary provides the app info
// @Description this endpoint provides the application version the current date and whether it is run on kubernetes
// @Accept json
// @Produce json
// @Success 200 {object} models.AppInfo
// @Router / [get]
func (h *AboutHandler) AppInfo(c *gin.Context) {
	c.JSON(http.StatusOK, &models.AppInfo{Version: version.Version, Date: time.Now().Unix(), Kubernetes: h.cfg.IsKubernetes})
}
