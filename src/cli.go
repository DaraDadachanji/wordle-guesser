package main

import (
	"flag"
	"fmt"
	"sort"
	"wordle/game"
)

type Flags struct {
	Guess          string
	Detailed       bool
	Rough          bool
	ShowAllGuesses bool
	Start          int
	End            int
	AnswersPerLine int
}

// start CLI application to suggest guesses
func RunGuesserCLI(showAllGuesses bool, answersPerLine int) {
	guesser := Guesser{Answers: &game.AllAnswers, Guesses: &game.AllGuesses}
	fmt.Printf("Suggested first guess: %s\n", firstGuess)
	for {
		guess := getUserInput("Guess: ")
		pattern := getUserInput("Pattern: ")
		if pattern == "ccccc" {
			fmt.Println("Congratulations!")
			return
		} else {
			guesser.GiveHint(guess, pattern)

			for i, answer := range *guesser.Answers {
				if i%answersPerLine == 0 {
					fmt.Print("\n")
				}
				fmt.Printf("%s ", answer)
			}
			if len(*guesser.Answers)%answersPerLine != 0 {
				fmt.Print("\n")
			}
			fmt.Print("\n")
			fmt.Printf("Narrowed to %d answers\n", len(*guesser.Answers))
			if len(*guesser.Answers) > 2 {
				var suggestions PairList
				if showAllGuesses {
					suggestions = guesser.AllGuesses()
				} else {
					suggestions = guesser.SuggestGuess()
				}
				sort.Sort(sort.Reverse(suggestions))
				for _, suggestion := range suggestions {
					averageOptions := float64(suggestion.Value) / float64(len(*guesser.Answers))
					fmt.Printf("Suggested Guess: %s, aggregate score %d, average remaining answers: %.2f\n", suggestion.Key, suggestion.Value, averageOptions)
				}
			}
		}
	}
}

func getFlags() Flags {
	guess := flag.String(
		"score",
		"",
		"a single guess to be scored",
	)
	detailed := flag.Bool(
		"detailed",
		false,
		"perform a detailed scoring of guesses from the rough list")
	showAllGuesses := flag.Bool(
		"allguesses",
		false,
		"show all guesses and scores after each hint instead of just the ones tied for best")
	start := flag.Int("start", 0, "starting point of for best guesses")
	end := flag.Int("end", 0, "ending point for best guesses")
	rough := flag.Bool("rough", false, "calculate alphabet and rough guess scores")
	answersPerLine := flag.Int("perLine", 5, "number of answers to print per line")
	flag.Parse()
	flags := Flags{
		Guess:          *guess,
		Detailed:       *detailed,
		Rough:          *rough,
		Start:          *start,
		End:            *end,
		ShowAllGuesses: *showAllGuesses,
		AnswersPerLine: *answersPerLine,
	}
	return flags
}

func getUserInput(prompt string) string {
	var userInput string
	for {
		fmt.Print(prompt)
		fmt.Scanln(&userInput)
		if len(userInput) == 5 {
			break
		} else {
			fmt.Printf("\nmust be exactly 5 letters\n")
		}
	}
	return userInput
}
