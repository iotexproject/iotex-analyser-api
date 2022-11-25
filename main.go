package main

import (
	"context"
	"embed"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/iotexproject/iotex-analyser-api/apiservice"
	"github.com/iotexproject/iotex-analyser-api/config"
	"github.com/iotexproject/iotex-analyser-api/db"
)

//go:embed templates
var templates embed.FS

//go:embed docs-html/*
var docsHtml embed.FS

const (
	ConfigPath = "ConfigPath"
)

func main() {
	configPath := os.Getenv(ConfigPath)
	//first load config from env
	if configPath == "" {
		configPath = config.FindDefaultConfigPath()
	}

	if configPath == "" {
		log.Fatalf("cannot determine default configuration path. %v, %v",
			config.DefaultConfigDirs,
			config.DefaultConfigFiles)
	}

	log.Printf("currently config path: %s", configPath)
	_, err := config.New(configPath)
	if err != nil {
		log.Fatalf("Failed to parse config: %v", err)
	}

	if config.Default.LogPath != "" {
		f, err := os.OpenFile(config.Default.LogPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Fatalf("Failed to open log file: %v", err)
		}
		log.SetOutput(io.MultiWriter(f, os.Stdout))
	}

	log.Printf("loaded config: %+v", config.Default)
	_, err = db.Connect()
	if err != nil {
		log.Fatalf("failed to connect DB, %v", err)
	}
	log.Printf("connected to DB")

	apiservice.DocsHTML = docsHtml

	ctx := context.Background()
	go func() {
		if err := apiservice.StartGRPCService(ctx); err != nil {
			log.Fatalf("failed to start GRPC API service, %v", err)
		}
	}()

	go func() {
		if err := apiservice.StartGRPCProxyService(templates); err != nil {
			log.Fatalf("failed to start HTTP API service, %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop
}
