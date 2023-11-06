package hangman

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"math/rand"
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

// PrintWord is a function that reveals a random set of letters in the word at the start of the game.
// It takes the target word as input and returns a string with some letters revealed (randomly chosen).
func PrintWord(word string) string {
    rand.Seed(time.Now().UnixNano()) // Seed the random number generator with the current time.

    // Calculate the number of letters to reveal (between 1 and len(word)/2 - 1)
    revealedCount := len(word)/2 - 1

    // Generate a random set of indices to reveal
    revealedIndices := make([]int, revealedCount)
    for i := 0; i < revealedCount; i++ {
        randomIndex := rand.Intn(len(word))
        revealedIndices[i] = randomIndex
    }

    var str string

    for i := 0; i < len(word); i++ {
        revealed := false
        for _, index := range revealedIndices {
            if i == index {
                str += string(word[i])
                revealed = true
                break
            }
        }
        if !revealed {
            str += "_"
        }
    }

    return str
}

// RevealLetters is a function responsible for revealing specific letters in the word.
// It takes the target word, a list of indices to reveal, and the current state of the revealed word.
// It updates the revealed word based on the provided indices and returns the updated revealed word.
func RevealLetters(word string, indices []int, revealedWord string) string {
    revealed := []rune(revealedWord) // Convert the revealed word to a rune slice for modification
    WordTab := []rune(word) // Convert the target word to a rune slice for access

    // Iterate through the provided indices and update the revealed word
    for _, index := range indices {
        if index >= 0 && index < len(WordTab) {
            revealed[index] = WordTab[index]
        }
    }

    return string(revealed) // Convert the updated revealed word back to a string
}

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

// Verify is a function that checks if a letter is present in the target word.
// It takes the target word and a letter as input and returns a slice of indices
// where the letter is found in the word. If the letter is not found, it returns nil.

func Verify(word, letter string) []int {
    WordTab := []rune(word)      // Convert the target word to a rune slice for character comparison
    RuneLetter := []rune(letter) // Convert the input letter to a rune slice for comparison
    var indices []int            // Initialize a slice to store indices where the letter is found

    // Iterate through the target word to find occurrences of the input letter
    for i := 0; i < len(WordTab); i++ {
        if RuneLetter[0] == WordTab[i] {
            indices = append(indices, i) // Add the index to the slice if the letter is found
        }
    }

    // If no occurrences of the letter are found, return nil
    if len(indices) == 0 {
        return nil
    }

    return indices
}

// WordList is a function that returns a random word from a text file or an error if any occurs.
// It takes the name of the text file as an argument and reads a list of words from the file.
// It then selects a random word from the list and returns it.

func WordList(textFile string) (string, error) {
    // Open the text file for reading
    file, err := os.Open(textFile)
    if err != nil {
        return "", err
    }
    defer file.Close()

    // Read the words from the file and store them in a slice
    var wordList []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        wordList = append(wordList, scanner.Text())
    }

    // Return an error if there's an issue while scanning the file
    if scanner.Err() != nil {
        return "", scanner.Err()
    }

    // Seed the random number generator with the current time
    rand.Seed(time.Now().UnixNano())

    // Select a random word from the list
    randomIndex := rand.Intn(len(wordList))
    randomWord := wordList[randomIndex]

    return randomWord, nil
}
