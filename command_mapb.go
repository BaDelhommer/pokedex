package main

import (
	"fmt"
)

func commandMapB(cfg *config, args ...string) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationAreaUrl)
	if err != nil {
		return err
	}

	if cfg.prevLocationAreaUrl == nil {
		return fmt.Errorf("you are on the first page")
	}

	for _, location := range resp.Results {
		fmt.Printf(" - %s\n", location.Name)
	}

	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaUrl = resp.Previous

	return nil
}
