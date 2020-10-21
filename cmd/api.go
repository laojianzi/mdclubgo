package main

import (
	"log"

	"github.com/laojianzi/mdclubgo/api"
)

func main() {
	if err := api.Server().Start(":3333"); err != nil {
		log.Fatal(err)
	}
}
