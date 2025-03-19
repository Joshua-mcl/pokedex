package main

import (
	"fmt"
	"os"
)

func callbackExit() {
	fmt.Printf("Closing the Pokedex... Goodbye!")
	os.Exit(0)
}
