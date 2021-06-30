package main

import (
	"go-scaffold/internal/app/router"
	"go-scaffold/internal/pkg/client"
	"go-scaffold/internal/pkg/config"
	"os"
)

func main() {
	config := config.NewConfig("insuremo-service").
		LoadLocalConf().
		OverrideWithRemoteConf()
	client := client.NewClient(config)
	r := router.SetupRouter(client)
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	} else {
		port = ":" + port
	}
	r.Run(port)
}
