package main

import (
	"strings"
	"wordle/game"
)

type Guesser struct {
	Answers *[]string
	Guesses *[]string
}

func (guesser *Guesser) GiveHint(guess string, pattern string) {
	hint := buildHint(guess, pattern)
	guesser.Answers = narrowList(guesser.Answers, hint)
	guesser.Guesses = narrowList(guesser.Guesses, hint)
}

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

func narrowList(list *[]string, hint game.Hint) *[]string {
	narrowedAnswers := []string{}
	for _, answer := range *list {
		if game.Validate(hint, answer) {
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
