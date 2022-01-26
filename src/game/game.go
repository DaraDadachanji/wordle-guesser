package game

import (
	"strings"
)

type LetterState int8

const (
	Absent LetterState = iota
	Present
	Correct
)

type HintPiece struct {
	Char         string
	State        LetterState
	AccountedFor bool
}

type Hint []HintPiece

type Letter struct {
	Char         string
	AccountedFor bool
}

type Word []Letter

func (word Word) IsPresent(char string) bool {
	for i, letter := range word {
		if !letter.AccountedFor {
			if letter.Char == char {
				word[i].AccountedFor = true
				return true
			}
		}
	}
	return false
}

// constructs a hint from a guess and an answer
func GetHint(guess string, answer string) Hint {
	_guess := build_word(guess)
	_answer := build_word(answer)
	hint := make(Hint, 5)
	for i := range _guess {
		if _guess[i].Char == _answer[i].Char {
			hint[i] = HintPiece{Char: _guess[i].Char, State: Correct, AccountedFor: false}
			_guess[i].AccountedFor = true
			_answer[i].AccountedFor = true
		}
	}
	for i, letter := range _guess {
		if !letter.AccountedFor {
			if _answer.IsPresent(letter.Char) {
				hint[i] = HintPiece{Char: letter.Char, State: Present, AccountedFor: false}
			} else {
				hint[i] = HintPiece{Char: letter.Char, State: Absent, AccountedFor: false}
			}
		}
	}
	return hint
}

// checks whether an answer is valid given a known hint
// optionally ignore correct (green) letters in the hint
func AnswerMatchesHint(hint Hint, answer string, ignoreCorrect bool) bool {
	_hint := make(Hint, len(hint))
	copy(_hint, hint) //mutable copy to work with
	_answer := build_word(answer)

	//check Correct
	for i, part := range _hint {
		matches := part.Char == _answer[i].Char
		shouldMatch := part.State == Correct
		if matches && shouldMatch {
			_hint[i].AccountedFor = true
			_answer[i].AccountedFor = true
			continue
		} else if matches != shouldMatch {
			//Don't filter out guesses where correct letters don't match
			//guessing them could provide valuable information
			if !ignoreCorrect {
				return false
			}
		} else if !matches && !shouldMatch {
			continue
		}
	}

	//check Present and Absent
	for _, part := range _hint {
		if !part.AccountedFor {
			if _answer.IsPresent(part.Char) {
				if part.State == Absent {
					return false
				}
			} else {
				if part.State == Present {
					return false
				}
			}
		}
	}
	return true
}

//init function for Word class
func build_word(str string) Word {
	chars := strings.Split(str, "")
	word := Word{}
	for _, char := range chars {
		letter := Letter{
			Char:         char,
			AccountedFor: false,
		}
		word = append(word, letter)
	}
	return word
}

// return true if the all letters in the hint are correct
func (hint Hint) AllCorrect() bool {
	for _, piece := range hint {
		if piece.State != Correct {
			return false
		}
	}
	return true
}
