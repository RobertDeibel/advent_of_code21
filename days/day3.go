package days

import (
	"log"
	"strconv"
)

type PuzzleD3 struct {
	bit_nums    []int64
	len_bit_num int
}

func (p *PuzzleD3) preprocess(args []string) {
	lines := readArgList(args)

	p.len_bit_num = len(lines[0])

	for _, line := range lines {
		num, err := strconv.ParseInt(line, 2, 0)

		if err != nil {
			log.Fatal(err)
		}

		p.bit_nums = append(p.bit_nums, num)
	}
}

func (p *PuzzleD3) calc_challenge() PuzzleReturn {
	var result PuzzleReturn

	gamma, epsilon := p.calc_gamma_rate()
	result.Args = append(result.Args, gamma, gamma, epsilon, epsilon, gamma*epsilon)

	result.Format_string = "Gamma rate binary: %012b base 10: %d, Epsilon rate binary: %012b, base 10: %d\n" +
		"multiplied: %d\n"

	oxygen := p.calc_oxygen_co(true)

	co := p.calc_oxygen_co(false)

	result.Args = append(result.Args, oxygen, oxygen, co, co, oxygen*co)

	result.Format_string += "Oxygen level: %012b base 10: %d, CO rate: %012b, base 10: %d\n" +
		"multiplied: %d\n"

	return result
}

func (p *PuzzleD3) calc_gamma_rate() (uint, uint) {
	buckets := map[int]int{}

	for _, num := range p.bit_nums {

		comp := func(shift int) bool {
			return comparator(num, shift)
		}

		for i := 0; i < p.len_bit_num; i++ {
			if comp(i) {
				buckets[i]++
			}
		}
	}

	var gamma_rate uint
	for i := 0; i < p.len_bit_num; i++ {
		if buckets[i] > (len(p.bit_nums) / 2) {
			gamma_rate += 1 << i
		}
	}

	epsilon_rate := (^gamma_rate) & ((1 << p.len_bit_num) - 1)

	return gamma_rate, epsilon_rate

}

func (p *PuzzleD3) calc_oxygen_co(calc_oxygen bool) int64 {
	var ones []int64
	var zeros []int64

	solution := p.bit_nums

	for i := p.len_bit_num - 1; i >= 0; i-- {
		one_count := 0

		comp := func(num int64) bool {
			return comparator(num, i)
		}

		for _, num := range solution {
			if comp(num) {
				one_count++
				ones = append(ones, num)

			} else {
				zeros = append(zeros, num)
			}
		}

		var take_ones bool
		if calc_oxygen {
			take_ones = len(ones) >= len(zeros)
		} else {
			take_ones = len(ones) < len(zeros)
		}

		if take_ones {
			solution = ones
		} else {
			solution = zeros
		}

		ones = []int64{}
		zeros = []int64{}

		if len(solution) == 1 {
			return solution[0]
		}

	}
	if len(solution) == 1 {
		return solution[0]
	} else {
		return -1
	}
}

func comparator(num int64, shift int) bool {
	return (num & (1 << shift)) > 0
}
