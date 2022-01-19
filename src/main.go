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
	flags := getFlags()
	if flags.Guess != "" {
		score := ScoreGuess(flags.Guess, &game.AllAnswers)
		fmt.Printf("%s: %d\n", flags.Guess, score)
	} else if flags.Rough {
		CalculateAlphabetValues()
		CalculateRoughGuessValues()
	} else if flags.Detailed {
		CalculateBestGuesses(flags.Start, flags.End)
	} else {
		HelpGuess()
	}

}

type Flags struct {
	Guess    string
	Detailed bool
	Rough    bool
	Start    int
	End      int
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
	start := flag.Int("start", 0, "starting point of for best guesses")
	end := flag.Int("end", 0, "ending point for best guesses")
	rough := flag.Bool("rough", false, "calculate alphabet and rough guess scores")
	flag.Parse()
	flags := Flags{
		Guess:    *guess,
		Detailed: *detailed,
		Rough:    *rough,
		Start:    *start,
		End:      *end,
	}
	return flags
}

func HelpGuess() {
	guesser := Guesser{Answers: &game.AllAnswers}
	fmt.Printf("Suggested first guess: %s\n", firstGuess)
	for {
		guess := getUserInput("Guess: ")
		pattern := getUserInput("Pattern: ")
		if pattern == "ccccc" {
			fmt.Println("Congratulations!")
			return
		} else {
			guesser.GiveHint(guess, pattern)
			fmt.Printf("Narrowed to %d answers\n", len(*guesser.Answers))
			for _, answer := range *guesser.Answers {
				fmt.Println(answer)
			}
			bestGuess := guesser.SuggestGuess()
			averageOptions := bestGuess.Value / len(*guesser.Answers)
			fmt.Printf("%s: %d (%d)\n", bestGuess.Key, bestGuess.Value, averageOptions)
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
