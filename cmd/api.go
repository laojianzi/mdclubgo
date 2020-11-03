package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/laojianzi/mdclubgo/api"
	"github.com/laojianzi/mdclubgo/conf"
	"github.com/laojianzi/mdclubgo/db"
	"github.com/laojianzi/mdclubgo/log"
)

func main() {
	if err := conf.Init(); err != nil {
		panic(fmt.Sprintf("Failed to initialize application: %v", err))
	}

	log.Init()
	db.Init()

	addr := fmt.Sprintf("%s:%s", conf.Server.HTTPAddr, conf.Server.HTTPPort)
	go func() {
		if err := api.Server().Start(addr); err != nil {
			log.Fatal("api start error: %s", err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c
	log.Info("Gracefully shutting down...")
	_ = api.Server().Shutdown()

	db.Close()
	log.Close()
}
