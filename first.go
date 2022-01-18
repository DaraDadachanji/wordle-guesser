package main

import (
	"encoding/json"
	"os"
	"wordle/game"
)

func CalculateBestGuesses(number int) {
	game.Initialize()
	guesses := GetTopRoughGuesses(number)
	scores := RankGuesses(guesses, &game.AllAnswers)
	output, _ := json.Marshal(scores)
	os.WriteFile("best_guesses.json", output, 0644)
}
