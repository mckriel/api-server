package redis

import "time"

type Session struct {
	ID         string    `json:"id" redis:"id"`
	User_ID    string    `json:"user_id" redis:"user_id"`
	Token      string    `json:"token" redis:"token"`
	IP_Address string    `json:"ip_address" redis:"ip_address"`
	User_Agent string    `json:"user_agent" redis:"user_agent"`
	Created_At time.Time `json:"created_at" redis:"created_at"`
	Expires_At time.Time `json:"expires_at" redis:"expires_at"`
	Active     bool      `json:"active" redis:"active"`
}