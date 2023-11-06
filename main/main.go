package main

import (
	"bufio"
	"fmt"
	"hangman" // Import a custom package called "hangman"
	"os"
	"strings"
	"time"
)

func main() {
	for {
		fmt.Print("\033[H\033[2J")                     // Clear the terminal screen
		hangman.Start("../start.txt")                  // Display the hangman starting image
		fmt.Print("\033[31m" + "INPUT : " + "\033[0m") // Display an input prompt

		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		// Remove leading/trailing whitespace from the user input
		input = strings.TrimSpace(input)

		if input == "99" {
			break // Exit the game if the user enters "99"
		} else if input == "" { // Start a new game if the user presses Enter without input
			// Load a random word from a file
			randomWord, err := hangman.WordList("../words.txt")
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			attempts := 10
			steps := 0
			usedFalse := ""
			usedTrue := ""

			fmt.Print("\033[H\033[2J")                      // Clear the terminal screen
			hangman.DisplayHangman("../hangman.txt", steps) // Display the initial hangman state
			revealedWord := hangman.PrintWord(randomWord)   // Initialize the revealed word
			fmt.Println(revealedWord)                       // Display the initial state of the word with underscores

			for attempts > 0 {
				input, err := hangman.Input() // Read user input (single letter)
				if err != nil {
					fmt.Println("Error:", err)
					return
				}

				indices := hangman.Verify(randomWord, input) // Check if the input letter exists in the word
				if indices != nil {
					fmt.Print("\033[H\033[2J")                                              // Clear the terminal screen
					revealedWord = hangman.RevealLetters(randomWord, indices, revealedWord) // Update the revealed word
					hangman.DisplayHangman("../hangman.txt", steps)                         // Display the hangman state
					fmt.Println("Letter found. Remaining attempts:", attempts)
					fmt.Println("Word:", revealedWord)
					usedTrue += input + " "
					fmt.Println("Used letter False:", "\033[31m"+usedFalse+"\033[0m")
					fmt.Println("Used letter True:", "\033[32m"+usedTrue+"\033[0m")
				} else {
					fmt.Print("\033[H\033[2J") // Clear the terminal screen
					attempts--
					steps++
					hangman.DisplayHangman("../hangman.txt", steps) // Display the updated hangman state
					fmt.Println("Letter not found. Remaining attempts:", attempts)
					fmt.Println("Word:", revealedWord)
					usedFalse += input + " "
					fmt.Println("Used letter False:", "\033[31m"+usedFalse+"\033[0m")
					fmt.Println("Used letter True:", "\033[32m"+usedTrue+"\033[0m")
				}

				if revealedWord == randomWord {
					fmt.Print("\033[H\033[2J")  // Clear the terminal screen
					hangman.Start("../win.txt") // Display a winning message
					fmt.Println("\033[33m"+"Congratulations! You guessed the word:", randomWord+"\033[0m")
					time.Sleep(5 * time.Second) // Sleep for 5 seconds
					break                       // Exit the game
				}
			}

			if revealedWord != randomWord {
				fmt.Print("\033[H\033[2J")    // Clear the terminal screen
				hangman.Start("../loose.txt") // Display a losing message
				fmt.Println("\033[31m"+"The word was:", randomWord+"\033[0m")
				time.Sleep(5 * time.Second) // Sleep for 5 seconds
			}
		}
	}
}
