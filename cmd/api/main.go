package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/mochivi/go-delivery-scheduler/internal/env"
	"github.com/mochivi/go-delivery-scheduler/internal/store"
)

func main() {

	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Create config
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	// Create store
	store := store.NewStorage(nil)

	// Create application
	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
