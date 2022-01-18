package main

import (
	"strings"
	"wordle/game"
)

type Guesser struct {
	Answers *[]string
}

func (guesser *Guesser) GiveHint(guess string, pattern string) {
	narrowedAnswers := []string{}
	hint := buildHint(guess, pattern)
	for _, answer := range *guesser.Answers {
		if game.Validate(hint, answer) {
			narrowedAnswers = append(narrowedAnswers, answer)
		}
	}
	guesser.Answers = &narrowedAnswers
}

func (guesser Guesser) SuggestGuess() *PairList {
	guesses := RankGuesses(guesser.Answers, guesser.Answers)
	return &guesses
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
