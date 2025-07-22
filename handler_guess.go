package main

func (cfg *config) handlerGuess(guess string, gyr [3]int) {

	charMap := []*[]string{
		&cfg.firstChar,
		&cfg.secondChar,
		&cfg.thirdChar,
		&cfg.fourthChar,
		&cfg.fifthChar,
	}

	checkedGuess, duplicate := duplicateLetterChecker(guess)

	if (gyr[0] != 0 && gyr[1] != 0) || duplicate { // letters can be anywhere, duplicate letters can be anywhere
		return
	}

	if gyr[0] != 0 && gyr[1] == 0 { // letters are where they are, or not in the word
		for i := range charMap {
			for j := range checkedGuess {
				if i != j {
					*charMap[i] = append(*charMap[i], string(checkedGuess[j]))
				}
			}
		}
		return
	}

	if gyr[0] == 0 && gyr[1] != 0 { // letters are not where they are, or not in the word
		for i := range charMap {
			for j := range checkedGuess {
				if i == j {
					*charMap[i] = append(*charMap[i], string(checkedGuess[j]))
				}
			}
		}
		return
	}

	if gyr[2] == 5 { // letters are not in the word
		for i := range charMap {
			for j := range checkedGuess {
				*charMap[i] = append(*charMap[i], string(checkedGuess[j]))
			}
		}
		return
	}
}

func duplicateLetterChecker(guess string) (string, bool) {
	counter := 0
	for _, chr := range guess {
		for i := range 5 {
			if string(guess[i]) == string(chr) {
				counter++
			}
		}
	}
	if counter != 5 { // duplicate letters can be anywhere
		return guess, true
	}
	return guess, false
}
