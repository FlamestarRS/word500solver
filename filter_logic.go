package main

import (
	"slices"
	"strings"

	"github.com/getlantern/words/american"
)

func filterLogic() []string {
	words := american.Words

	fiveLetterWords := []string{}
	for _, word := range words {
		if strings.Contains(word, "'") {
			continue
		}

		if len(word) == 5 {
			fiveLetterWords = append(fiveLetterWords, word)
		}
	}

	fiveLetterWordsNoDuplicates := []string{}
	for _, word := range fiveLetterWords {
		if word[0] == word[1] || word[0] == word[2] || word[0] == word[3] || word[0] == word[4] {
			continue
		}
		if word[1] == word[2] || word[1] == word[3] || word[1] == word[4] {
			continue
		}
		if word[2] == word[3] || word[2] == word[4] {
			continue
		}
		if word[3] == word[4] {
			continue
		}

		fiveLetterWordsNoDuplicates = append(fiveLetterWordsNoDuplicates, word)
	}

	return fiveLetterWordsNoDuplicates

	// firstChar := []string{"a", "b"} // example list, real list comes from guess info

}

func (cfg *config) repeatFiltering() []string {
	newPossibleWords := []string{}
	for _, word := range cfg.possibleWords {
		skip := false

		for guess, gyr := range cfg.guessMap {
			guessChars := []rune(guess)
			wordChars := []rune(word)
			numValidChars := (gyr[0] + gyr[1])
			counter := 0

			for i := range 5 {
				if slices.Contains(wordChars, guessChars[i]) {
					counter++
				}
			}

			if counter < numValidChars {
				skip = true
				continue
			}
		}

		if skip {
			continue
		}

		_, guessed := cfg.guessMap[word]
		if !slices.Contains(cfg.firstChar, string(word[0])) &&
			!slices.Contains(cfg.secondChar, string(word[1])) &&
			!slices.Contains(cfg.thirdChar, string(word[2])) &&
			!slices.Contains(cfg.fourthChar, string(word[3])) &&
			!slices.Contains(cfg.fifthChar, string(word[4])) &&
			!slices.Contains(newPossibleWords, word) &&
			!guessed {
			newPossibleWords = append(newPossibleWords, word)
		}
	}

	return newPossibleWords
}
