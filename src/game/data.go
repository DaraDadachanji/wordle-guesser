package game

import (
	"bufio"
	"bytes"
	_ "embed"
	"io"
)

var (
	AllGuesses []string
	AllAnswers []string

	//go:embed guesses.txt
	guesses []byte

	//go:embed answers.txt
	answers []byte
)

func Initialize() {
	AllGuesses, _ = readLines(bytes.NewReader(guesses))
	AllAnswers, _ = readLines(bytes.NewReader(answers))
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(file io.Reader) ([]string, error) {

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
