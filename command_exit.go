package main

import (
	"fmt"
	"os"
)

func callbackExit() error {
	fmt.Printf("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
