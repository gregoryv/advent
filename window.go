package advent

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func IncreasingWindow(filename string, win int) {
	in := loadInts(filename)
	IncreasingWindowTo(os.Stdout, in, win)
}

func IncreasingWindowTo(w io.Writer, in []int, win int) {
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
	fmt.Fprintln(w, count)
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
