package days

import (
	"log"
	"strconv"
	"strings"
)

type PuzzleD5 struct {
	lines []pline
	max_x int
	max_y int
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

		if xs > p.max_x {
			p.max_x = xs
		}

		if err != nil {
			log.Fatal(err)
		}

		ye, err := strconv.Atoi(split[len(split)-1])

		if ye > p.max_y {
			p.max_y = ye
		}

		if err != nil {
			log.Fatal(err)
		}

		strings.Split(split[1], " ")

		ys, err := strconv.Atoi(split[0])

		if ys > p.max_y {
			p.max_y = ys
		}

		if err != nil {
			log.Fatal(err)
		}

		xe, err := strconv.Atoi(split[len(split)-1])

		if xe > p.max_x {
			p.max_x = xe
		}

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

	floor := make([]int, p.max_x*p.max_y)

	return result
}
