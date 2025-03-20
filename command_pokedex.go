package main

import (
	"fmt"
)

func callbackPokedex(cfg *config, args ...string) error {

	//pokemon, ok := cfg.caughtPokemon[pokemonName]
	fmt.Println("Pokemon in Pokedex:")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", pokemon.Name)

	}

	return nil
}
