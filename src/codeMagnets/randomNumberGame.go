package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	seconds := time.Now().Unix() // get current date and time as integer
	rand.Seed(seconds)           // seed the random number generator
	randomNumber := rand.Intn(100) + 1
	fmt.Println(randomNumber)
	var userGuess int
	for guessChances := 1; guessChances <= 10; guessChances++ {

		fmt.Print("Guess the Number: ")
		fmt.Scanln(&userGuess)

		if randomNumber > userGuess {
			fmt.Println("Guess is Low")
		} else if randomNumber < userGuess {
			fmt.Println("Guess is High")
		} else {
			fmt.Println("You guessed it right")
			return
		}
	}

	fmt.Println("Sorry out of chances")
}
