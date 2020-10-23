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
	if err := api.Server().Start(":8080"); err != nil {
		log.Fatal("api start error: %s", err.Error())
	}

	log.Close()
}
