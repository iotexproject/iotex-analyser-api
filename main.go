package main

import (
	"context"
	"embed"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/iotexproject/iotex-analyser-api/apiservice"
	"github.com/iotexproject/iotex-analyser-api/common/tasks"
	"github.com/iotexproject/iotex-analyser-api/config"
	"github.com/iotexproject/iotex-analyser-api/db"
	"github.com/iotexproject/iotex-core/pkg/routine"
	"go.uber.org/zap"
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

	task := routine.NewRecurringTask(tasks.ChainSyncWorker, time.Hour*6)
	if err := task.Start(ctx); err != nil {
		log.Fatal("Failed to start chainsync routine.", zap.Error(err))
	}
	defer func() {
		if err := task.Stop(ctx); err != nil {
			log.Fatal("Failed to stop chainsync routine.", zap.Error(err))
		}
	}()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop
}
