package main

import (
	"encoding/json"
	"fmt"
	"os"
	"wordle/game"
)

func CalculateBestGuesses(start int, end int) {
	if start > end {
		panic("start must be < end")
	}
	game.Initialize()
	guesses := GetRoughGuesses(start, end)
	scores := RankGuesses(guesses, &game.AllAnswers)
	output, _ := json.Marshal(scores)
	filename := fmt.Sprintf("best_guesses_%d-%d.json", start, end)
	os.WriteFile(filename, output, 0644)
}
