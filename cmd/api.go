package main

import (
	"github.com/laojianzi/mdclubgo/api"
	"github.com/laojianzi/mdclubgo/log"
)

func main() {
	log.Init()

	if err := api.Server().Start(":3333"); err != nil {
		log.Fatal("api start error: %s", err.Error())
	}

	log.Close()
}
