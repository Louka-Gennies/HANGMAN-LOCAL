package hangman

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
