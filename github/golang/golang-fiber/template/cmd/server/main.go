package main

import (
	"log"

	"golang-fiber-template/internal/server"
)

func main() {
	// Inicia o servidor Fiber
	s := server.NewServer()
	if err := s.Listen(":3000"); err != nil {
		log.Fatalf("Falha ao iniciar o servidor: %v", err)
	}
}
