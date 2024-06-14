package main

import (
	"fmt"
	"log"
)

func commandMap(cfg *config) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Location Areas:")

	for _, location := range resp.Results {
		fmt.Printf(" - %s\n", location.Name)
	}

	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaUrl = resp.Previous

	return nil
}
