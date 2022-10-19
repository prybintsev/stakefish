package api

import (
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

type AboutResponse struct {
	Version    string `json:"version"`
	Date       int64  `json:"date"`
	Kubernetes bool   `json:"kubernetes"`
}

func (h *AboutHandler) About(c *gin.Context) {
	c.JSON(http.StatusOK, &AboutResponse{Version: version.Version, Date: time.Now().Unix(), Kubernetes: h.cfg.IsKubernetes})
}
