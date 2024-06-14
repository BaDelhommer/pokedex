package main

import "github.com/BaDelhommer/pokedex/internal/pokeapi"

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaUrl *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
	}
	startRepl(&cfg)
}
