package days

import (
	"log"
	"strconv"
	"strings"
)

type PuzzleD2 struct {
	commands []command
}

type command struct {
	direction string
	amount    int
}

type ship struct {
	horizontal int
	depth      int
	aim        int
}

func (p *PuzzleD2) preprocess(args []string) {
	lines := readArgList(args)

	for _, line := range lines {
		split_str := strings.Split(line, " ")

		if len(split_str) < 2 {
			log.Fatal("Error in command chain")
		}

		amount, err := strconv.Atoi(split_str[1])

		if err != nil {
			log.Fatal(err)
		}

		command := command{
			direction: split_str[0],
			amount:    amount,
		}

		p.commands = append(p.commands, command)
	}

}

func (p *PuzzleD2) calc_challenge() PuzzleReturn {

	result := PuzzleReturn{}
	ship1 := ship{}

	ship1.traverse(p.commands)
	result.Args = append(result.Args, ship1.horizontal, ship1.depth, ship1.horizontal*ship1.depth)

	ship2 := ship{}

	ship2.traverseWithAim(p.commands)

	result.Args = append(result.Args, ship2.horizontal, ship2.depth, ship2.horizontal*ship2.depth)

	result.Format_string = "final coordinates first try:\n x: %d y: %d\nmultiplied: %d\n" +
		"final coordinates second try:\n x: %d y: %d\nmultiplied: %d\n"

	return result
}

func (s *ship) traverse(commands []command) {
	for _, c := range commands {
		switch c.direction {
		case "forward":
			s.horizontal += c.amount
		case "up":
			s.depth -= c.amount
		case "down":
			s.depth += c.amount
		}
	}
}

func (s *ship) traverseWithAim(commands []command) {
	for _, c := range commands {
		switch c.direction {
		case "forward":
			s.horizontal += c.amount
			s.depth += s.aim * c.amount
		case "up":
			s.aim -= c.amount
		case "down":
			s.aim += c.amount
		}
	}
}
