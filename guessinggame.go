// Guessing game
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Heard you're a psychic...prove it to me!")

	// random number generator
	source := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(source)
	secretNumber := randomizer.Intn(10) // generates numbers from 0 to 10 inclusive

	var guess int
	var sum int
	// prompt user for number until game is won
	for {
		fmt.Println("What's the magic number?")
		fmt.Scan(&guess)
		if guess > secretNumber {
			fmt.Println("That's bigger")
		} else if guess < secretNumber {
			fmt.Println("That's smaller")
		} else {
			fmt.Println("Awesome!")
			break
		}
		sum++
	}
	fmt.Println("You got it on trial", sum+1)
}
