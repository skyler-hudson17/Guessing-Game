package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	// Generating Random Number
	rand.Seed(time.Now().UnixNano())
	var min int = 0
	var max int = 10
	var randNum int = rand.Intn(max-min+1) - min

	// Variables for the Game
	var guess string
	var attempts int = 0

	// Game Loop
	fmt.Println("================== Welcome to the Guessing Game ==================")
	for {
		fmt.Print("Please input your guess: ")
		fmt.Scan(&guess)

		// Converting the random number to a string
		randNumStr := strconv.Itoa(randNum)

		// Guessing
		if guess != randNumStr {
			fmt.Print("Please try again: ")
			attempts++
			continue
		} else if guess == randNumStr {
			fmt.Println("You guessed the number in:", attempts, "attempts")
			break
		}
	}
}
