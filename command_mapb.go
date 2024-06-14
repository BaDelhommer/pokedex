package main

import (
	"fmt"
	"log"
)

func commandMapB(cfg *config) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationAreaUrl)
	if err != nil {
		log.Fatal(err)
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
