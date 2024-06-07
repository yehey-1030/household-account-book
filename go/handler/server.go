package handler

import (
	"golang.org/x/sys/unix"
	"log"
	"net/http"
	"os"
	"os/signal"
)

type Server struct {
	*rootRouter
	address string
}

func NewServer(router *rootRouter, address string) *Server {
	return &Server{router, address}
}

func (s *Server) Start() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, unix.SIGTERM)

	go func() {
		httpServer := &http.Server{
			Addr:    s.address,
			Handler: s.rootRouter,
		}
		log.Fatal(httpServer.ListenAndServe())
	}()
	<-signalChan
}
