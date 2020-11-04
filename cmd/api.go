package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/laojianzi/mdclubgo/api"
	"github.com/laojianzi/mdclubgo/conf"
	"github.com/laojianzi/mdclubgo/db"
	"github.com/laojianzi/mdclubgo/log"
)

func main() {
	if err := conf.Init(); err != nil {
		panic(fmt.Sprintf("failed to initialize application: %v", err))
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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	if err := api.Server().Shutdown(ctx); err != nil {
		log.Fatal("api server shutdown error: %s", err.Error())
	}
	cancel()

	db.Close()
	log.Close()
}
