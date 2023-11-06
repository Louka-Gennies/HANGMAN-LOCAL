package hangman

import (
	"bufio"
	"os"
	"fmt"
)

// DisplayHangman displays the hangman ASCII art from a file based on the specified range of lines.
// It uses ANSI escape codes to color the text blue for a visually appealing hangman display.
func DisplayHangman(filename string, attempts int) error {
    // Open the file containing the hangman ASCII art.
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    // Create a scanner to read the file line by line and store each line in a slice.
    scanner := bufio.NewScanner(file)
    lines := make([]string, 0)

    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    // Return an error if there was an issue while scanning the file.
    if scanner.Err() != nil {
        return scanner.Err()
    }

    // Calculate the range of lines to display based on the number of incorrect attempts.
    // Each incorrect attempt typically adds 7 lines to the hangman display.
    startLine := attempts * 7
    endLine := startLine + 7

    // Ensure that the start and end lines are within the bounds of the available lines.
    if startLine < 0 {
        startLine = 0
    }
    if endLine > len(lines) {
        endLine = len(lines)
    }

    // Display the selected lines in blue color using ANSI escape codes.
    for i := startLine; i < endLine; i++ {
        fmt.Println("\033[34m" + lines[i] + "\033[0m")
    }

    // Return nil to indicate that the function executed successfully.
    return nil
}
