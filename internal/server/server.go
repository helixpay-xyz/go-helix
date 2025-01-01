package server

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/helixpay-xyz/go-helix/internal/api"
	"github.com/helixpay-xyz/go-helix/internal/scan"
)

type Server struct {
	apiHandler *api.APIHandler
	scanner    *scan.Scanner
}

func NewServer() *Server {
	return &Server{
		apiHandler: api.NewAPIHandler(),
		scanner:    scan.NewScanner(),
	}
}

func (s *Server) Run() {
	go func() {
		s.apiHandler.Run()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	log.Println("Shutting down server...")
}
