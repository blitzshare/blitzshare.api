package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"blitzshare.fileshare.api/app/dependencies"
)

type Server struct {
	httpServer *http.Server
	closeFn    func()
}

func (s *Server) Stop() error {
	err := s.httpServer.Shutdown(context.Background())
	s.closeFn()
	return err
}

func start(router http.Handler, deps *dependencies.Dependencies, closeFn func()) (Server, error) {
	s := Server{
		closeFn: closeFn,
	}

	s.httpServer = &http.Server{
		Addr:         fmt.Sprintf(":%v", deps.Config.Server.Port),
		Handler:      router,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	go func() {
		err := s.httpServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			s.closeFn()
			log.Fatalf("error on http server: %v\n", err)
		}
	}()

	return s, nil
}

func Start(router http.Handler, deps *dependencies.Dependencies, wg *sync.WaitGroup) Server {
	wg.Add(1)
	s, err := start(router, deps, func() {
		wg.Done()
	})

	if err != nil {
		log.Fatalf("error on http server: %v\n", err)
	}
	return s
}
