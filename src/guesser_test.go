package main

import (
	"testing"
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
	if score != 2 {
		t.Error()
	}
}
