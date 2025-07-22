package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	startRepl()
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := config{
		guessMap:      map[string][3]int{},
		inputError:    "Please enter your guess and results in the following format: trick 014",
		firstChar:     []string{"j", "q", "x", "z"},
		secondChar:    []string{"j", "q", "x", "z"},
		thirdChar:     []string{"j", "q", "x", "z"},
		fourthChar:    []string{"j", "q", "x", "z"},
		fifthChar:     []string{"j", "q", "x", "z"},
		possibleWords: filterLogic(),
	}
	fmt.Println("Welcome to word500solver!")
	for {
		fmt.Print("Enter your 5 letter word and results: ")
		scanner.Scan()
		guess := cleanInput(scanner.Text())
		if len(guess) <= 0 {
			continue
		}

		if strings.HasPrefix(guess[0], "!") {
			fmt.Println(cfg.possibleWords)
			continue
		}

		if len(guess) != 2 || len(guess[0]) != 5 || len(guess[1]) != 3 {
			fmt.Println(cfg.inputError)
			continue
		}

		if guess[1] == "500" {
			fmt.Println("Congratulations! Press enter to exit...")
			scanner.Scan()
			break
		}

		gyr, err := gyrArray(guess[1])
		if err != nil {
			fmt.Println(err)
			continue
		}

		cfg.guessMap[guess[0]] = gyr

		printMap(cfg)
		cfg.handlerGuess(guess[0], cfg.guessMap[guess[0]])
		cfg.possibleWords = cfg.repeatFiltering()
		fmt.Println("Possible Solutions:", len(cfg.possibleWords), "[Enter '!' to view]")

	}
}

type config struct {
	guessMap      map[string][3]int
	inputError    string
	firstChar     []string
	secondChar    []string
	thirdChar     []string
	fourthChar    []string
	fifthChar     []string
	possibleWords []string
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	split := strings.Fields(lowered)
	return split
}

func gyrArray(gyr string) ([3]int, error) {
	cleanGyr := [3]int{}
	for i, chr := range gyr {
		cleanGyr[i] = int(chr - '0')
	}

	if err := gyrCheck(cleanGyr); err != nil {
		return [3]int{}, err
	}

	return cleanGyr, nil
}

func gyrCheck(gyr [3]int) error {
	sum := 0
	for _, num := range gyr {
		sum += num
	}
	if sum != 5 {
		return fmt.Errorf("results must add up to 5")
	}
	return nil
}

func printMap(cfg config) error {
	if len(cfg.guessMap) <= 0 {
		fmt.Println("no guesses yet")
	}

	fmt.Printf("%v guesses so far:\n", len(cfg.guessMap))
	for word, gyr := range cfg.guessMap {
		fmt.Println(word, gyr)
	}

	return nil
}
