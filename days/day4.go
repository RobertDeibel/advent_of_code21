package days

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type PuzzleD4 struct {
	boards []*bingoB
	nums   []int
	winner winningPair
}

type bingoB struct {
	fields   [25]*field
	selected bool
}

type field struct {
	num    int
	marked bool
}

type winningPair struct {
	sum       int
	drawn_num int
}

func (p *PuzzleD4) preprocess(args []string) {
	lines := readArgList(args)

	board := &bingoB{}
	b := 0
	for i, line := range lines {
		if i == 0 {
			split := strings.Split(line, ",")
			for _, c := range split {
				num, err := strconv.Atoi(c)

				if err != nil {
					log.Fatal(err)
				}
				p.nums = append(p.nums, num)
			}

		} else {
			s := strings.TrimSpace(line)
			if len(s) > 0 {
				split := strings.Split(s, " ")
				for _, c := range split {

					if len(c) <= 0 {
						continue
					}

					num, err := strconv.Atoi(c)
					if err != nil {
						log.Fatal(err)
					}

					board.fields[b] = &field{num: num}
					b++
				}
				if b == 25 {
					b = 0
					p.boards = append(p.boards, board)
					board = &bingoB{}
				}
			}
		}
	}
}

func (p *PuzzleD4) calc_challenge() PuzzleReturn {

	result := PuzzleReturn{}

	for !p.bingo() {
	}

	sum := p.winner.sum

	result.Args = append(result.Args, sum, p.winner.drawn_num, sum*p.winner.drawn_num)

	result.Format_string = "Sum of unmarked fields of winner is: %d, winning draw: %d,\n" +
		"multiplied: %d\n"

	for len(p.nums) > 0 {
		p.bingo()
	}

	sum = p.winner.sum

	result.Args = append(result.Args, sum, p.winner.drawn_num, sum*p.winner.drawn_num)

	result.Format_string += "Last Winner\n" + result.Format_string

	return result
}

func (p *PuzzleD4) bingo() bool {
	// pop first
	next_num := p.nums[0]
	p.nums = p.nums[1:]

	winning_round := false

	for _, board := range p.boards {
		for _, f := range board.fields {
			if f.num == next_num {
				f.marked = true

			}
		}
	}
	for _, board := range p.boards {

		if board.checkBingo() && !board.selected {
			p.winner = winningPair{sum: board.sum_unmarked(), drawn_num: next_num}
			board.selected = true

			winning_round = true
		}
	}

	return winning_round

}

func (board *bingoB) checkBingo() bool {
	rows := map[int]int{}
	cols := map[int]int{}

	for i, f := range board.fields {
		if f.marked {
			rows[i/5]++
			cols[i%5]++
			if rows[i/5] == 5 || cols[i%5] == 5 {
				return true
			}
		}
	}

	return false
}

func (board *bingoB) sum_unmarked() int {
	sum := 0
	for _, f := range board.fields {
		if !f.marked {
			sum += f.num
		}
	}

	return sum
}

func (board bingoB) String() string {
	s := ""
	m := ""
	var nums []interface{}
	for i, f := range board.fields {
		s += "%2d"
		if f.marked {
			m += "XX"
		} else {
			m += "00"
		}
		if i%5 == 4 {
			s += "\n"
			m += "\n"
		} else {
			s += " "
			m += " "
		}
		nums = append(nums, f.num)
	}

	return fmt.Sprintf(s+"\n"+m, nums...)
}
