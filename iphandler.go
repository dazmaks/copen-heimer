package main

import (
	"bytes"
	"log"

	"github.com/iverly/go-mcping/mcping"
	"gorm.io/gorm"
)

func IPHandler(db *gorm.DB, addr string) {
	pinger := mcping.NewPinger()

	//log.Printf("Info [mcstatus]: Pinging %s\n", ip)
	response, err := pinger.Ping(addr, 25565)
	if err != nil {
		log.Printf("Error [mcstatus]: %s", err)
	} else {
		var playerlist bytes.Buffer
		for i := 0; i < len(response.Sample); i++ {
			playerlist.WriteString(response.Sample[i].Name)
			playerlist.WriteByte(',')
		}
		go WriteServer(db, Server{
			IP:       addr,
			Latency:  response.Latency,
			Motd:     response.Motd,
			Version:  response.Version,
			Protocol: response.Protocol,
			Online:   response.PlayerCount.Online,
			Max:      response.PlayerCount.Max,
			Players:  playerlist.String(),
		})
	}
}
