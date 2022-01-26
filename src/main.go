package main

import (
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
		HelpGuess(flags.ShowAllGuesses)
	}

}

func HelpGuess(showAllGuesses bool) {
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

			for _, answer := range *guesser.Answers {
				fmt.Println(answer)
			}
			fmt.Printf("Narrowed to %d answers\n", len(*guesser.Answers))
			if len(*guesser.Answers) > 2 {
				suggestions := guesser.SuggestGuess(showAllGuesses)
				sort.Sort(sort.Reverse(suggestions))
				for _, suggestion := range suggestions {
					averageOptions := float64(suggestion.Value) / float64(len(*guesser.Answers))
					fmt.Printf("Suggested Guess: %s, aggregate score %d, average remaining answers: %.2f\n", suggestion.Key, suggestion.Value, averageOptions)
				}
			}
		}
	}
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
		hint := game.GetHint(guess, answer)
		hints = append(hints, hint)
	}
	score := 0
	for _, hint := range hints {
		if hint.IsExactMatch() {
			continue
		}
		for _, answer := range *answers {
			if game.AnswerMatchesHint(hint, answer, false) {
				score++
			}
		}
	}
	return score
}
