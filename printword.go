package hangman

import (
	"math/rand"
	"time"
)

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
