package main

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/BaDelhommer/pokedex/internal/pokeapi"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}

	pokemonName := args[0]

	resp, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	fmt.Printf("Throwing a pokeball at %v...\n", resp.Name)

	catchChance := rand.Intn(resp.BaseExperience)
	if catchChance < 50 {
		fmt.Printf("%v was caught!\n", resp.Name)
		pokeapi.PokedexMap[resp.Name] = resp
	} else {
		fmt.Printf("%v escaped!", resp.Name)
	}

	return nil
}
