package main

import (
	"fmt"
	"log"

	"github.com/242617/pace/server"
	"github.com/242617/pace/version"
)

func main() {
	log.SetFlags(log.Lshortfile)
	fmt.Printf("%s started\n", version.Application)
	log.Fatal(server.Init())
}
