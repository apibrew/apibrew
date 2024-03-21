package rest

import (
	"time"
)

type HealthResponse struct {
	Status string `json:"status"`
}

type AuthenticationRequest struct {
	Username      string `json:"username"`
	Password      string `json:"password"`
	Term          string `json:"term"`
	MinimizeToken bool   `json:"minimize"`
}

type Token struct {
	Term       string `json:"term"`
	Content    string `json:"content"`
	Expiration time.Time
}

type AuthenticationResponse struct {
	Token Token `json:"token"`
}

type RefreshTokenRequest struct {
	Token string `json:"token"`
	Term  string `json:"term"`
}

type RefreshTokenResponse struct {
	Token Token `json:"token"`
}
