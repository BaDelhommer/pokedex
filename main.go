package main

import (
	"time"

	"github.com/BaDelhommer/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaUrl *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
	}
	startRepl(&cfg)
}
