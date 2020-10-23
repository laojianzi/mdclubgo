package main

import (
	"fmt"

	"github.com/laojianzi/mdclubgo/api"
	"github.com/laojianzi/mdclubgo/conf"
	"github.com/laojianzi/mdclubgo/log"
)

func main() {
	if err := conf.Init(); err != nil {
		panic(fmt.Sprintf("Failed to initialize application: %v", err))
	}

	log.Init()

	addr := fmt.Sprintf("%s:%s", conf.Server.HTTPAddr, conf.Server.HTTPPort)
	if err := api.Server().Start(addr); err != nil {
		log.Fatal("api start error: %s", err.Error())
	}

	log.Close()
}
