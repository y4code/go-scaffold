package client

import (
	"go-scaffold/internal/pkg/config"
)

type Client struct {
	Config *config.Config
}

func NewClient(config *config.Config) *Client {
	return &Client{Config: config}
}
