package main

import (
	"errors"
	"fmt"
	"math/rand"
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

	const threshHold = 50
	catchChance := rand.Intn(resp.BaseExperience)
	if catchChance < threshHold {
		fmt.Printf("%v was caught!\n", resp.Name)
		cfg.caughtPokemon[resp.Name] = resp
	} else {
		return fmt.Errorf("%v escaped", resp.Name)
	}

	return nil
}
