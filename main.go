package main

import (
	"log"
	"os"

	"cave/api"
	"cave/config"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	dirname, err := os.UserHomeDir()
    if err != nil {
        log.Fatal( err )
    }
	log.Printf("Home Dir: %v\n", dirname + "/cave/.env")

	config := config.NewConfig(dirname + "/cave/.env")
	api.ConfigAndRunApp(config)
}
