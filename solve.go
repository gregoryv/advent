package aoc2021

func Day1_part1(in []int) int {
	var count int
	prev := in[0]
	// skip first value
	for i := 1; i < len(in); i++ {
		if in[i] > prev {
			count++
		}
		// always save previous
		prev = in[i]
	}
	return count
}
