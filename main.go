package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/iotexproject/iotex-analyser-api/apiservice"
	"github.com/iotexproject/iotex-analyser-api/config"
	"github.com/iotexproject/iotex-analyser-api/db"
)

func main() {
	configPath := config.FindDefaultConfigPath()
	if configPath == "" {
		log.Fatalf("Cannot determine default configuration path. %v, %v",
			config.DefaultConfigDirs,
			config.DefaultConfigFiles)
	}

	log.Printf("Currently config path: %s", configPath)
	_, err := config.New(configPath)
	if err != nil {
		log.Fatalf("Failed to parse config: %v", err)
	}
	_, err = db.Connect()
	if err != nil {
		log.Fatalf("failed to connect DB, %v", err)
	}
	ctx := context.Background()
	go func() {
		if err := apiservice.StartGRPCService(ctx); err != nil {
			log.Fatalf("failed to start GRPC API service, %v", err)
		}
	}()

	go func() {
		if err := apiservice.StartGRPCProxyService(); err != nil {
			log.Fatalf("failed to start HTTP API service, %v", err)
		}
	}()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop
}
