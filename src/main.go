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
		averageOptions := float64(score) / float64(len(game.AllAnswers))
		fmt.Printf("%s: %d (average remaining answers: %.2f)\n", flags.Guess, score, averageOptions)
	} else if flags.Rough {
		CalculateAlphabetValues()
		CalculateRoughGuessValues()
	} else if flags.Detailed {
		CalculateBestGuesses(flags.Start, flags.End)
	} else {
		RunGuesserCLI(flags.ShowAllGuesses)
	}

}

// start CLI application to suggest guesses
func RunGuesserCLI(showAllGuesses bool) {
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

// score and sort a list of guesses based on a list of answers
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

// score a single guess and write the result to a channel
func WriteScoreAsync(guess string, answers *[]string, scores chan Pair) {
	score := ScoreGuess(guess, answers)
	scores <- Pair{Key: guess, Value: score}
}

// generate hints for all answers in a list,
// then count how many answers are still valid given each hint
// return the combined total remaining answers for all cases
func ScoreGuess(guess string, answers *[]string) int {
	hints := []game.Hint{}
	for _, answer := range *answers {
		hint := game.GetHint(guess, answer)
		hints = append(hints, hint)
	}
	score := 0
	for _, hint := range hints {
		if hint.AllCorrect() {
			continue
			// this causes exact matches to add 0 instead of 1
			// useful as a tiebreaker between valid answers and guesses
		}
		for _, answer := range *answers {
			if game.AnswerMatchesHint(hint, answer, false) {
				score++
			}
		}
	}
	return score
}
