package main

import (
	"github.com/y4code/go-scaffold/internal/app/router"
	"github.com/y4code/go-scaffold/internal/app/service"
	"github.com/y4code/go-scaffold/internal/pkg/client"
	"github.com/y4code/go-scaffold/internal/pkg/config"
	"github.com/y4code/go-scaffold/internal/pkg/storage"
)

func main() {

	// TODO: refactor with Viper
	config := config.NewConfig("service-name").
		LoadLocalConf().
		OverrideWithRemoteConf()

	memory := storage.NewStorage(storage.TypeMemory)
	exampleService := service.NewExampleService(memory)
	marketService := service.NewMarketService(memory)

	client := client.NewClient(config, exampleService, marketService)

	r := router.SetupRouter(client)

	r.Run(":8080")
}
