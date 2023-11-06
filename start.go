package hangman

import (
	"bufio"
	"os"
	"fmt"
)

// Start function is responsible for displaying the initial hangman or game-related content
// from a specified file. It uses ANSI escape codes to apply red color for a visual effect.
// It takes the name of the file containing the content to display as an argument.

func Start(filename string) error {
    file, err := os.Open(filename) // Open the specified file.
    if err != nil {
        return err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    lines := make([]string, 0)

    // Read the content of the file line by line and store each line in a slice.
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    // Return an error if there's an issue while scanning the file.
    if scanner.Err() != nil {
        return scanner.Err()
    }

    // Display the first 16 lines of the content using red color (ANSI escape codes).
    for i := 0; i < 16; i++ {
        fmt.Println("\033[31m" + lines[i] + "\033[0m")
    }

    return nil
}
