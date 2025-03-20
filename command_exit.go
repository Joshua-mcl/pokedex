package main

import (
	"fmt"
	"os"
)

func callbackExit(cfg *config) error {
	fmt.Printf("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
