package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/prybintsev/stakefish/internal/iptools"
)

type IPLookupHandler struct {
	logEntry *logrus.Entry
}

func NewIPLookupHandler(logEntry *logrus.Entry) IPLookupHandler {
	return IPLookupHandler{logEntry: logEntry}
}

type LookupResponse struct {
	ClientIP  string    `json:"client_ip"`
	CreatedAt int64     `json:"created_at"`
	Domain    string    `json:"domain"`
	Addresses []Address `json:"addresses"`
}

type Address struct {
	IP string `json:"ip"`
}

func (h *IPLookupHandler) Lookup(c *gin.Context) {
	domain := c.Query("domain")
	if domain == "" {
		writeErrorResponse(c, http.StatusBadRequest, "missing domain parameter")
		return
	}

	ips, err := iptools.GetIpv4sByDomain(domain)
	if err != nil {
		writeErrorResponse(c, http.StatusInternalServerError, "failed to lookup the given domain")
		return
	}

	var addresses []Address
	for _, ip := range ips {
		addresses = append(addresses, Address{IP: ip.String()})
	}
	res := LookupResponse{
		ClientIP:  c.ClientIP(),
		CreatedAt: time.Now().Unix(),
		Domain:    domain,
		Addresses: addresses,
	}
	c.JSON(http.StatusOK, res)
}

type ValidateRequest struct {
	IP string `json:"ip"`
}

type ValidateResponse struct {
	Status bool `json:"status"`
}

func (h *IPLookupHandler) Validate(c *gin.Context) {
	var req ValidateRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		writeErrorResponse(c, http.StatusBadRequest, "malformed request")
	}

	c.JSON(http.StatusOK, ValidateResponse{Status: iptools.IsValidIPv4(req.IP)})
}
