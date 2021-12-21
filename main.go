package main

import (
	"flag"

	days "advent_of_code_21/days"
)

func main() {
	dayPtr := flag.Int("day", 1, "Challenge number")

	flag.Parse()

	var day days.Puzzle
	switch *dayPtr {
	case 1:
		day = days.Puzzle{Challenge: &days.PuzzleD1{}, Day: 1}
	case 2:
		day = days.Puzzle{Challenge: &days.PuzzleD2{}, Day: 2}
	case 3:
		day = days.Puzzle{Challenge: &days.PuzzleD3{}, Day: 3}
	case 4:
		day = days.Puzzle{Challenge: &days.PuzzleD4{}, Day: 4}
	case 5:
		day = days.Puzzle{Challenge: &days.PuzzleD5{}, Day: 5}
	}
	days.DoChallenge(flag.Args(), day)
}
