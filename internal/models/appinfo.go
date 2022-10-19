package models

type AppInfo struct {
	Version    string `json:"version"`
	Date       int64  `json:"date"`
	Kubernetes bool   `json:"kubernetes"`
}
