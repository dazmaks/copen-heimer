package main

import (
	"log"

	"gorm.io/gorm"
)

type Server struct {
	IP       string
	Latency  uint
	Motd     string
	Version  string
	Protocol int
	Online   int
	Max      int
	Players  string
}

func WriteServer(db *gorm.DB, server Server) {
	log.Printf("Info [database]: Writing %s\n", server.IP)
	db.Create(server)
}
