package hangman

import (
	"os"
    "bufio"
    "math/rand"
    "time"
)

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
