package main

import (
	"log"

	"ipfs-visualizer/config"
	"ipfs-visualizer/internal/app"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("config error: %v", err)
	}

	a, err := app.NewApp(cfg)
	if err != nil {
		log.Fatalf("app init error: %v", err)
	}

	if err := a.Start(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
