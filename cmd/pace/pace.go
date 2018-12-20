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
	fmt.Printf("%s starting\n", version.Application)

	err := storage.Init()
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(server.Init())
}
