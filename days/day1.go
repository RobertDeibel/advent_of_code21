package days

import (
	"log"
	"strconv"
)

type PuzzleD1 struct {
	nums []int
}

// preprocess takes command line arguments and prepares the data struct PuzzleD1
func (p *PuzzleD1) preprocess(args []string) {

	lines := readArgList(args)

	for _, line := range lines {
		num, err := strconv.Atoi(line)

		if err != nil {
			log.Fatal(err)
		}

		p.nums = append(p.nums, num)
	}

}

// calc_challenge calculates the result of the challenge for the day
func (p *PuzzleD1) calc_challenge() PuzzleReturn {

	result := PuzzleReturn{}

	result.Args = append(result.Args, p.getNumIncrease())

	result.Args = append(result.Args, p.getSlidingWindowIncrease())

	result.Format_string = "per line increases: %d\nper 3 sliding window increases: %d\n"

	return result
}

func (p *PuzzleD1) getNumIncrease() int {
	numIncreases := 0

	last := p.nums[0]
	for _, num := range p.nums {
		if num > last {
			numIncreases++
		}

		last = num
	}

	return numIncreases
}

func (p *PuzzleD1) getSlidingWindowIncrease() int {
	numIncreases := 0

	sumSlice := func(s []int) int {
		sum := 0
		for _, num := range s {
			sum += num
		}

		return sum
	}

	last := sumSlice(p.nums[:4])

	for i := 1; i < len(p.nums); i++ {
		next := sumSlice(p.nums[i : i+3])

		if next > last {
			numIncreases++
		}
		last = next
	}

	return numIncreases

}
