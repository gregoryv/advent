package advent

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func Example_IncreaseInts_input() {
	fmt.Print(IncreasingWindow(loadInts("testdata/1.input"), 1))
	// output: 1288
}

func Example_IncreaseInts_given() {
	fmt.Print(IncreasingWindow(loadInts("testdata/1.given"), 1))
	// output: 7
}

func Example_IncreaseWindow_input() {
	fmt.Print(IncreasingWindow(loadInts("testdata/1.input"), 3))
	// output: 1311
}

func Example_IncreaseWindow_given() {
	fmt.Print(IncreasingWindow(loadInts("testdata/1.given"), 3))
	// output: 5
}

func IncreasingWindow(in []int, win int) int {
	sum := func(in []int, w int) int {
		var s int
		for i := 0; i < win; i++ {
			s += in[w-i]
		}
		return s
	}

	var count int
	prev := sum(in, win)

	// skip first value, nothing to compare with
	for i := win - 1; i < len(in); i++ {
		s := sum(in, i)
		if s > prev {
			count++
		}
		// always save previous
		prev = s
	}
	return count
}

// ----------------------------------------

func loadInts(filename string) []int {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(data), "\n")
	in := make([]int, len(lines))
	for i, line := range lines {
		if line == "" { // skip empty lines
			continue
		}
		in[i], err = strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
	}
	return in
}
