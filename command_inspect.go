package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("no pokemon name provided")
	}

	pokemonName := args[0]

	pokemon, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		return errors.New("you have not caught this pokemon yet")
	}

	fmt.Printf("-Name: %v\n", pokemon.Name)
	fmt.Printf("-Height: %v\n", pokemon.Height)
	fmt.Printf("-Weight: %v\n", pokemon.Weight)
	for _, stat := range pokemon.Stats {
		fmt.Printf("-%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	for _, typ := range pokemon.Types {
		fmt.Printf("-Type: %v\n", typ.Type.Name)
	}

	for _, ability := range pokemon.Abilities {
		fmt.Printf("-Ability: %v\n", ability.Ability.Name)
	}

	return nil
}
