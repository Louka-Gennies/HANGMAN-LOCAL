package hangman

import (
	"strings"
	"bufio"
	"os"
	"fmt"
)

// Input function is responsible for taking input from the user, specifically a single letter.
// It returns the user's input as a string in uppercase for consistency.
// It ensures that the input is a valid single letter and prompts the user for input until a valid letter is provided.
func Input() (string, error) {
    for {
        fmt.Print("Enter a single letter: ") // Prompt the user for input

        // Read input from the standard input (keyboard)
        reader := bufio.NewReader(os.Stdin)
        input, err := reader.ReadString('\n')
        if err != nil {
            return "", err
        }

        // Remove leading and trailing whitespace and convert the input to uppercase for consistency
        input = strings.TrimSpace(strings.ToUpper(input))

        // Check if the input is a valid single letter (A to Z)
        if len(input) == 1 && input >= "A" && input <= "Z" {
            return input, nil // Return the valid input
        }

        fmt.Println("Invalid input. Please enter a single letter.") // Display an error message for invalid input
    }
}
