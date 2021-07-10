package client

import (
	"go-scaffold/internal/app/service"
	"go-scaffold/internal/pkg/config"
)

type Client struct {
	Config         *config.Config
	ExampleService service.ExampleService
	MarketService  service.MarketService
}

func NewClient(config *config.Config, exampleService service.ExampleService, marketService service.MarketService) *Client {
	return &Client{Config: config, ExampleService: exampleService, MarketService: marketService}
}
