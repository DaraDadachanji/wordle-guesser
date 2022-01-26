package main

import (
	"flag"
	"fmt"
)

type Flags struct {
	Guess          string
	Detailed       bool
	Rough          bool
	ShowAllGuesses bool
	Start          int
	End            int
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
	flag.Parse()
	flags := Flags{
		Guess:          *guess,
		Detailed:       *detailed,
		Rough:          *rough,
		Start:          *start,
		End:            *end,
		ShowAllGuesses: *showAllGuesses,
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
