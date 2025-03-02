package main

import ("fmt"
		"os"
		"bufio"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print(" >")

	scanner.Scan()
	text := scanner.Text()

	fmt.Println("echoing: ", text)

	
}