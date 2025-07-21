package main

func (cfg *config) handlerGuess(guess string, gyr [3]int) {

	if gyr[0] != 0 && gyr[1] != 0 { // letters can be anywhere
		return
	}
	if gyr[0] != 0 && gyr[1] == 0 { // letters are where they are, or not in the word
		cfg.firstChar = append(cfg.firstChar, string(guess[1]))
		cfg.firstChar = append(cfg.firstChar, string(guess[2]))
		cfg.firstChar = append(cfg.firstChar, string(guess[3]))
		cfg.firstChar = append(cfg.firstChar, string(guess[4]))
		cfg.secondChar = append(cfg.secondChar, string(guess[0]))
		cfg.secondChar = append(cfg.secondChar, string(guess[2]))
		cfg.secondChar = append(cfg.secondChar, string(guess[3]))
		cfg.secondChar = append(cfg.secondChar, string(guess[4]))
		cfg.thirdChar = append(cfg.thirdChar, string(guess[0]))
		cfg.thirdChar = append(cfg.thirdChar, string(guess[1]))
		cfg.thirdChar = append(cfg.thirdChar, string(guess[3]))
		cfg.thirdChar = append(cfg.thirdChar, string(guess[4]))
		cfg.fourthChar = append(cfg.fourthChar, string(guess[0]))
		cfg.fourthChar = append(cfg.fourthChar, string(guess[1]))
		cfg.fourthChar = append(cfg.fourthChar, string(guess[2]))
		cfg.fourthChar = append(cfg.fourthChar, string(guess[4]))
		cfg.fifthChar = append(cfg.fifthChar, string(guess[0]))
		cfg.fifthChar = append(cfg.fifthChar, string(guess[1]))
		cfg.fifthChar = append(cfg.fifthChar, string(guess[2]))
		cfg.fifthChar = append(cfg.fifthChar, string(guess[3]))
		return
	}

	if gyr[0] == 0 && gyr[1] != 0 { // letters are not where they are, or not in the word
		cfg.firstChar = append(cfg.firstChar, string(guess[0]))
		cfg.secondChar = append(cfg.secondChar, string(guess[1]))
		cfg.thirdChar = append(cfg.thirdChar, string(guess[2]))
		cfg.fourthChar = append(cfg.fourthChar, string(guess[3]))
		cfg.fifthChar = append(cfg.fifthChar, string(guess[3]))
		return
	}

	if gyr[2] == 5 { // letters are not in the word
		cfg.firstChar = append(cfg.firstChar, string(guess[0]))
		cfg.firstChar = append(cfg.firstChar, string(guess[1]))
		cfg.firstChar = append(cfg.firstChar, string(guess[2]))
		cfg.firstChar = append(cfg.firstChar, string(guess[3]))
		cfg.firstChar = append(cfg.firstChar, string(guess[4]))
		cfg.secondChar = append(cfg.secondChar, string(guess[0]))
		cfg.secondChar = append(cfg.secondChar, string(guess[1]))
		cfg.secondChar = append(cfg.secondChar, string(guess[2]))
		cfg.secondChar = append(cfg.secondChar, string(guess[3]))
		cfg.secondChar = append(cfg.secondChar, string(guess[4]))
		cfg.thirdChar = append(cfg.thirdChar, string(guess[0]))
		cfg.thirdChar = append(cfg.thirdChar, string(guess[1]))
		cfg.thirdChar = append(cfg.thirdChar, string(guess[2]))
		cfg.thirdChar = append(cfg.thirdChar, string(guess[3]))
		cfg.thirdChar = append(cfg.thirdChar, string(guess[4]))
		cfg.fourthChar = append(cfg.fourthChar, string(guess[0]))
		cfg.fourthChar = append(cfg.fourthChar, string(guess[1]))
		cfg.fourthChar = append(cfg.fourthChar, string(guess[2]))
		cfg.fourthChar = append(cfg.fourthChar, string(guess[3]))
		cfg.fourthChar = append(cfg.fourthChar, string(guess[4]))
		cfg.fifthChar = append(cfg.fifthChar, string(guess[0]))
		cfg.fifthChar = append(cfg.fifthChar, string(guess[1]))
		cfg.fifthChar = append(cfg.fifthChar, string(guess[2]))
		cfg.fifthChar = append(cfg.fifthChar, string(guess[3]))
		cfg.fifthChar = append(cfg.fifthChar, string(guess[4]))
		return

	}
}
