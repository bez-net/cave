package main

import (
	"log"

	"cave/api"
	"cave/config"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	config := config.NewConfig()
	api.ConfigAndRunApp(config)
}
