package main

import (
	"log"

	"github.com/kartikx04/gopher/internal/env"
	"github.com/kartikx04/gopher/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":3000"),
	}

	store := store.NewStorage(nil)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
