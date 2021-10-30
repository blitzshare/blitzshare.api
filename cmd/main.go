package main

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"blitzshare.fileshare.api/app/config"
	"blitzshare.fileshare.api/app/dependencies"
	"blitzshare.fileshare.api/app/server"
	log "github.com/sirupsen/logrus"
)

func initLog() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	// log.SetLevel(log.WarnLevel)
}

func main() {
	cfg, err := config.Load()
	initLog()
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
