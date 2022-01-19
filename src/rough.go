package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"absurdle/game"
)

func GetRoughGuesses(start int, end int) *[]string {
	var guessValues PairList
	jsonFile, _ := os.Open("top_guesses.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &guessValues)
	guesses := []string{}
	for _, guess := range guessValues {
		guesses = append(guesses, guess.Key)
	}
	if start >= 0 {
		shortlist := guesses[start:end]
		return &shortlist
	} else {
		return &guesses
	}

}

func CalculateRoughGuessValues() {
	guesses := game.AllGuesses
	var alphabet map[string]int
	jsonFile, _ := os.Open("alphabet.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &alphabet)
	guessValues := map[string]int{}
	for _, guess := range guesses {
		value := 0
		chars := map[string]bool{}
		for _, char := range strings.Split(guess, "") {
			chars[char] = true
		}
		for char := range chars {
			value += alphabet[char]
		}
		guessValues[guess] = value
	}
	topGuesses := sortMapByValue(guessValues, false)
	output, _ := json.Marshal(topGuesses)
	os.WriteFile("top_guesses.json", output, 0644)
}

func CalculateAlphabetValues() {
	answers := game.AllAnswers
	alphabet := map[string]int{}
	for _, letter := range strings.Split("abcdefghijklmnopqrstuvwxyz", "") {
		value := 0
		for _, answer := range answers {
			if strings.Contains(answer, letter) {
				value++
			}
		}
		alphabet[string(letter)] = value
	}
	for char, value := range alphabet {
		fmt.Printf("%s: %d\n", char, value)
	}
	output, _ := json.Marshal(alphabet)
	os.WriteFile("alphabet.json", output, 0644)

}
