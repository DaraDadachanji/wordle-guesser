package main

import (
	"testing"
	"wordle/game"
)

func TestScoring(t *testing.T) {
	answers := []string{"again", "guava"}
	scores := PairList{}
	for _, guess := range answers {
		score := ScoreGuess(guess, &answers)
		scores = append(scores, Pair{Key: guess, Value: score})
	}
	if scores[0].Value != scores[1].Value {
		t.Error()
	}
}

func TestGuava(t *testing.T) {
	answers := []string{"again", "guava"}
	score := ScoreGuess("guava", &answers)
	if score != 1 {
		t.Error()
	}
}

func TestDonor(t *testing.T) {
	hint := buildHint("roate", "ppaaa")
	answer := "donor"
	valid := game.AnswerMatchesHint(hint, answer, false)
	if valid {
		t.Error()
	}
}

func TestGroup(t *testing.T) {
	hint := buildHint("prior", "ccapa")
	answer := "group"
	valid := game.AnswerMatchesHint(hint, answer, false)
	if valid {
		t.Error()
	}
}
