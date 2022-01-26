package main

import (
	"strings"
	"wordle/game"
)

type Guesser struct {
	Answers *[]string
	Guesses *[]string
}

// filter answer and guess lists given a guess and pattern
func (guesser *Guesser) GiveHint(guess string, pattern string) {
	hint := buildHint(guess, pattern)
	guesser.Answers = narrowList(guesser.Answers, hint, false)
	guesser.Guesses = narrowList(guesser.Guesses, hint, true)
}

// rank remaining guesses based on remaining answers
// return guesses tied for best score
func (guesser Guesser) SuggestGuess() PairList {
	allGuesses := RankGuesses(guesser.Guesses, guesser.Answers)
	bestScore := allGuesses[0].Value
	bestGuesses := PairList{}
	for _, guess := range allGuesses {
		if guess.Value > bestScore {
			return bestGuesses
		} else {
			bestGuesses = append(bestGuesses, guess)
		}
	}
	return allGuesses //all guesses are tied
}

func (guesser Guesser) AllGuesses() PairList {
	allGuesses := RankGuesses(guesser.Guesses, guesser.Answers)
	return allGuesses
}

func narrowList(list *[]string, hint game.Hint, ignoreCorrect bool) *[]string {
	narrowedAnswers := []string{}
	for _, answer := range *list {
		if game.AnswerMatchesHint(hint, answer, ignoreCorrect) {
			narrowedAnswers = append(narrowedAnswers, answer)
		}
	}
	return &narrowedAnswers
}

func buildHint(guess string, pattern string) game.Hint {
	hint := game.Hint{}
	states := strings.Split(pattern, "")
	for i, char := range strings.Split(guess, "") {
		var state game.LetterState
		if states[i] == "c" {
			state = game.Correct
		} else if states[i] == "p" {
			state = game.Present
		} else if states[i] == "a" {
			state = game.Absent
		}
		part := game.HintPiece{Char: char, State: state}
		hint = append(hint, part)
	}
	return hint
}
