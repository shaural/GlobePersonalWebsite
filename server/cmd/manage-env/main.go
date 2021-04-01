package main

import (
	"context"
	"flag"
	"log"

	"github.com/shaural/GlobePersonalWebsite/server/pkg/db"
)

func main() {
	initialize := flag.Bool("initialize", false, "set to initialize postgres db")
	flag.Parse()

	if *initialize {
		log.SetPrefix("manage-env::initialize ")
		loadDB, err := db.NewDatabase(context.Background())
		if err != nil {
			log.Fatalf("Unable to establish connection to postgres database Error: %v", err)
		}
		defer loadDB.Close()

		if err = loadDB.Initialize(); err != nil {
			log.Fatalf("Encountered error initializing. Rolled back any changes. Error: %v", err)
		}
	}
}
