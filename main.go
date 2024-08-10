package main

import (
	"log"

	"github.com/MiguelCock/MiguelCock-st0263.git/pserver"
)

func main() {
	err := pserver.Serv()
	if err != nil {
		log.Fatal(err)
	}
}
