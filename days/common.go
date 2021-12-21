package days

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readArgList(args []string) []string {
	path := args[0]

	puzzle, err := readInputList(path)

	if err != nil {
		log.Fatal(err)
	}

	return puzzle
}

func readInputList(path string) ([]string, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	fs := bufio.NewScanner(f)

	var lines []string
	for fs.Scan() {
		line := fs.Text()

		lines = append(lines, line)
	}

	return lines, err

}

// DoChallange calculates the puzzle and returns the result
func DoChallenge(args []string, p Puzzle) {
	fmt.Printf("Day %d\n", p.Day)

	p.preprocess(args)

	p.output = p.Challenge.calc_challenge()

	fmt.Printf(p.output.Format_string, p.output.Args...)
}

// Challenge represents the calculations on a challenge
type Challenge interface {
	preprocess(args []string)
	calc_challenge() PuzzleReturn
}

// Puzzle is the base structure of a day in the advent of code
type Puzzle struct {
	Challenge
	Day    int
	output PuzzleReturn
}

// PuzzleReturn is the return structure of the puzzles and used for output
type PuzzleReturn struct {
	Format_string string
	Args          []interface{}
}

func wait() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter: ")
	reader.ReadString('\n')
}
