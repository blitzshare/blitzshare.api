package main

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"blitzshare.api/app/config"
	"blitzshare.api/app/dependencies"
	"blitzshare.api/app/server"
	log "github.com/sirupsen/logrus"
)

func initLog() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	initLog()
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config %v\n", err)
	}
	deps, err := dependencies.NewDependencies(cfg)
	if err != nil {
		log.Fatalf("failed to load dependencies %v\n", err)
	}
	router := server.NewRouter(deps)
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)

	wg := &sync.WaitGroup{}

	httpServer := server.Start(router, deps, wg)
	log.Printf("server running on port %d", cfg.Server.Port)
	<-signals

	err = httpServer.Stop()
	if err != nil {
		log.Fatalf("failed to stop http server %v", err)
	}
}
