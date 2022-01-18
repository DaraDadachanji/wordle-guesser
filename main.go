package main

import (
	"flag"
	"fmt"
	"sort"
	"wordle/game"
)

const firstGuess = "roate"

func main() {
	game.Initialize()
	guess := flag.String(
		"score",
		"",
		"a guess to be evaluated",
	)
	doBestGuesses := flag.Int("best", 0, "number of top rough guesses to score")
	doRoughGuesses := flag.Bool("rough", false, "calculate alphabet and rough guess scores")
	flag.Parse()
	if *guess != "" {
		score := ScoreGuess(*guess, &game.AllAnswers)
		fmt.Printf("%s: %d\n", *guess, score)
	} else if *doRoughGuesses {
		CalculateAlphabetValues()
		CalculateRoughGuessValues()
	} else if *doBestGuesses > 0 {
		CalculateBestGuesses(*doBestGuesses)
	} else {
		HelpGuess()
	}

}

func HelpGuess() {
	guesser := Guesser{Answers: &game.AllAnswers}
	fmt.Printf("Suggested first guess: %s\n", firstGuess)
	for {
		guess := getUserInput("Guess: ")
		pattern := getUserInput("Pattern: ")
		guesser.GiveHint(guess, pattern)
		fmt.Printf("Narrowed to %d answers\n", len(*guesser.Answers))
		options := guesser.SuggestGuess()
		for _, option := range *options {
			fmt.Printf("%s: %d\n", option.Key, option.Value)
		}
	}
}

func getUserInput(prompt string) string {
	var userInput string
	fmt.Print(prompt)
	fmt.Scanln(&userInput)
	return userInput
}

func RankGuesses(guesses *[]string, answers *[]string) PairList {
	scoreChannel := make(chan Pair, 50)
	for _, guess := range *guesses {
		go WriteScoreAsync(guess, answers, scoreChannel)
	}
	scores := PairList{}
	for range *guesses {
		scores = append(scores, <-scoreChannel)
	}
	sort.Sort(scores)
	return scores
}

func WriteScoreAsync(guess string, answers *[]string, scores chan Pair) {
	score := ScoreGuess(guess, answers)
	scores <- Pair{Key: guess, Value: score}
}

func ScoreGuess(guess string, answers *[]string) int {
	hints := []game.Hint{}
	for _, answer := range *answers {
		hint := game.Check(guess, answer)
		hints = append(hints, hint)
	}
	score := 0
	for _, hint := range hints {
		for _, answer := range *answers {
			if game.Validate(hint, answer) {
				score++
			}
		}
	}
	return score
}
