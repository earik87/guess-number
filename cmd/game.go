/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// gameCmd represents the game command
var gameCmd = &cobra.Command{
	Use:   "game",
	Short: "Start the number guessing game",
	Long: `A fun CLI number guessing game where you try to guess
a random number between 1 and 100.`,
	Run: runGame,
}

func runGame(cmd *cobra.Command, args []string) {
	fmt.Println("\nWelcome to the Number Guessing Game!")
	fmt.Println("Please select the difficulty level:")
	fmt.Println("1. Easy (10 chances)")
	fmt.Println("2. Medium (5 chances)")
	fmt.Println("3. Hard (3 chances)")
	fmt.Print("\nEnter your choice (1-3): ")

	reader := bufio.NewReader(os.Stdin)
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	// Add these type definitions at the top of the file after imports
	type difficulty struct {
		name    string
		chances int
	}

	// Add this map as a package-level variable
	var difficultyLevels = map[string]difficulty{
		"1": {"Easy", 10},
		"2": {"Medium", 5},
		"3": {"Hard", 3},
	}

	// Then replace the switch statement in runGame with:
	var chances int
	if level, exists := difficultyLevels[choice]; exists {
		chances = level.chances
		fmt.Printf("\nGreat! You have selected the %s difficulty level!\n", level.name)
	} else {
		chances = 5
		fmt.Println("\nInvalid choice. Defaulting to Medium mode.")
	}

	target := rand.Intn(100) + 1
	fmt.Printf("\nI'm thinking of a number between 1 and 100. You have %d chances.\n", chances)

	for i := 0; i < chances; i++ {
		fmt.Printf("\nChances left: %d\n", chances-i)
		fmt.Print("Enter your guess: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Please enter a valid number!")
			i-- // Don't count invalid inputs
			continue
		}

		if guess == target {
			fmt.Printf("\nCongratulations! You've guessed the number in %d attempts!\n", i+1)
			return
		} else if guess < target {
			fmt.Println("Too low! Try a higher number.")
		} else {
			fmt.Println("Too high! Try a lower number.")
		}
	}

	fmt.Printf("\nGame Over! The number was %d\n", target)
}

func init() {
	rootCmd.AddCommand(gameCmd)
}
