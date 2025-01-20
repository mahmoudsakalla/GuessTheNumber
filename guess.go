package main

import (
	"fmt"
	"math/rand"
)

func main() {
	target := rand.Intn(100) + 1 // Generate a number between 1 and 100
	var guess int

	fmt.Println("Welcome to the Guess the Number game!")
	fmt.Println("I'm thinking of a number between 1 and 100. Can you guess it?")

	for {
		fmt.Print("Enter your guess: ")
		fmt.Scanln(&guess)

		if guess < target {
			fmt.Println("Too low!")
		} else if guess > target {
			fmt.Println("Too high!")
		} else {
			fmt.Println("Congratulations! You guessed the number!")
			break
		}
	}
}
