package game

import (
	"bufio"
	"os"
)

var (
	AllGuesses []string
	AllAnswers []string
)

func Initialize() {
	AllGuesses, _ = readLines("guesses.txt")
	AllAnswers, _ = readLines("answers.txt")
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
