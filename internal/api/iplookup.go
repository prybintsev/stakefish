package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/prybintsev/stakefish/internal/db/lookup"
	"github.com/prybintsev/stakefish/internal/iptools"
	"github.com/prybintsev/stakefish/internal/models"
)

type IPLookupHandler struct {
	logEntry   *logrus.Entry
	lookupRepo *lookup.Repo
}

func NewIPLookupHandler(logEntry *logrus.Entry, lookupRepo *lookup.Repo) IPLookupHandler {
	return IPLookupHandler{logEntry: logEntry, lookupRepo: lookupRepo}
}

func (h *IPLookupHandler) Lookup(c *gin.Context) {
	domain := c.Query("domain")
	if domain == "" {
		writeErrorResponse(c, http.StatusBadRequest, "missing domain parameter")
		return
	}

	var msg string
	ips, err := iptools.GetIpv4sByDomain(domain)
	if err != nil {
		msg = "failed to lookup the given domain"
		h.logEntry.WithError(err).WithField("domain", domain).Error(msg)
		writeErrorResponse(c, http.StatusBadRequest, msg)
		return
	}

	var addresses []models.Address
	for _, ip := range ips {
		addresses = append(addresses, models.Address{IP: ip.String()})
	}
	res := models.Lookup{
		ClientIP:  c.ClientIP(),
		CreatedAt: time.Now().Unix(),
		Domain:    domain,
		Addresses: addresses,
	}

	err = h.lookupRepo.InsertLookup(c, res)
	if err != nil {
		msg = "failed to save lookup history"
		h.logEntry.WithError(err).Error(msg)
		writeErrorResponse(c, http.StatusInternalServerError, msg)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *IPLookupHandler) Validate(c *gin.Context) {
	var req models.ValidateRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		writeErrorResponse(c, http.StatusBadRequest, "malformed request")
		return
	}

	c.JSON(http.StatusOK, models.ValidateResponse{Status: iptools.IsValidIPv4(req.IP)})
}

func (h *IPLookupHandler) History(c *gin.Context) {
	lookups, err := h.lookupRepo.GetLastLookups(c)
	if err != nil {
		writeErrorResponse(c, http.StatusInternalServerError, "failed to retrieve lookups history")
		return
	}

	c.JSON(http.StatusOK, lookups)
}
