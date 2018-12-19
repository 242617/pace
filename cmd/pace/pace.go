package main

import (
	"fmt"
	"log"

	"github.com/242617/pace/server"
	"github.com/242617/pace/storage"
	"github.com/242617/pace/version"
)

func main() {
	log.SetFlags(log.Lshortfile)

	err := storage.Init()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s started\n", version.Application)
	log.Fatal(server.Init())
}
