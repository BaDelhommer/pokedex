package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) == 0 {
		return errors.New("pokedex is empty")
	}

	if len(args) == 0 {
		fmt.Println("Pokemon in Pokedex: ")
		for name := range cfg.caughtPokemon {
			fmt.Printf("- %v\n", name)
		}
		return nil
	}

	if len(args) == 1 {
		pokemonName := args[0]
		pokemon, ok := cfg.caughtPokemon[pokemonName]
		if !ok {
			return fmt.Errorf("you have not caught %v yet", pokemonName)
		}
		fmt.Printf("You have %v in your pokedex\n", pokemon.Name)
	}

	return nil
}
