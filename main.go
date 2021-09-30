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
        log.Printf( "ERROR Dirname: %v\n", err )
    }

	config := config.NewConfig(dirname + "/cave/.env")
	api.ConfigAndRunApp(config)
}
