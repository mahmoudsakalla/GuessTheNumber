package main

import (
	"fmt"
	"math/rand"
)

//for the random number generation

func main() {
	target := rand.Intn(100) + 1 // this will generate a number between 1-100
	var guess int

	fmt.Println("Welcome to the Guess the Number game!!!")
	fmt.Println("The computer is thinking of a number between 1 and 100. Can you guess it?")

	//The program will then enter a loop where it tells the user if the guess is high or low
	//Once correclty guessed, the program will end and print a congrats message.

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
