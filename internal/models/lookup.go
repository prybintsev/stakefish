package models

type Lookup struct {
	ClientIP  string    `json:"client_ip"`
	CreatedAt int64     `json:"created_at"`
	Domain    string    `json:"domain"`
	Addresses []Address `json:"addresses"`
}

type Address struct {
	IP string `json:"ip"`
}

type ValidateRequest struct {
	IP string `json:"ip"`
}

type ValidateResponse struct {
	Status bool `json:"status"`
}
