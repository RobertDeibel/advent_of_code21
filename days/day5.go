package days

import (
	"log"
	"strconv"
	"strings"
)

type PuzzleD5 struct {
	lines []pline
}

type pline struct {
	start coordinate
	end   coordinate
}

type coordinate struct {
	x int
	y int
}

func (p *PuzzleD5) preprocess(args []string) {
	lines := readArgList(args)

	for _, line := range lines {
		split := strings.Split(line, ",")
		xs, err := strconv.Atoi(split[0])

		if err != nil {
			log.Fatal(err)
		}

		ye, err := strconv.Atoi(split[len(split)-1])

		if err != nil {
			log.Fatal(err)
		}

		strings.Split(split[1], " ")

		ys, err := strconv.Atoi(split[0])

		if err != nil {
			log.Fatal(err)
		}

		xe, err := strconv.Atoi(split[len(split)-1])

		if err != nil {
			log.Fatal(err)
		}

		start := coordinate{x: xs, y: ys}
		end := coordinate{x: xe, y: ye}

		p.lines = append(p.lines, pline{start, end})

	}
}

func (p *PuzzleD5) calc_challenge() PuzzleReturn {
	result := PuzzleReturn{}

	return result
}
