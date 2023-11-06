package hangman

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
