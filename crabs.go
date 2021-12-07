package advent

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"os"
	"strings"
)

func HorizontalAlign(filename string) {
	fh, err := os.Open(filename)
	shouldNot(err)
	defer fh.Close()
	HorizontalAlignTo(os.Stdout, fh)
}

func HorizontalAlignTo(w io.Writer, r io.Reader) {
	data, _ := ioutil.ReadAll(r)
	values := strings.Split(string(data), ",")
	positions := toInts(values)

	var lowest int
	for _, pos := range positions {
		cost := alignTo(positions, pos, CostDirect)
		if lowest == 0 || cost < lowest {
			lowest = cost
		}
	}
	fmt.Fprintln(w, lowest)
}

// alignTo returns cost to align input to the given value
func alignTo(input []int, v int, alg Cost) int {
	var sum int
	for _, in := range input {
		sum += alg(in, v)
	}
	return sum
}

type Cost func(in, v int) int

func CostDirect(in, v int) int {
	return int(math.Abs(float64(in) - float64(v)))
}

// ----------------------------------------

func HorizontalAlignReal(filename string) {
	fh, err := os.Open(filename)
	shouldNot(err)
	defer fh.Close()
	HorizontalAlignRealTo(os.Stdout, fh)
}

func HorizontalAlignRealTo(w io.Writer, r io.Reader) {
	data, _ := ioutil.ReadAll(r)
	values := strings.Split(string(data), ",")
	positions := toInts(values)

	var max int
	for _, pos := range positions {
		if pos > max {
			max = pos
		}
	}

	var lowest int
	for pos := 1; pos <= max; pos++ {
		cost := alignTo(positions, pos, CostReal)
		if lowest == 0 || cost < lowest {
			lowest = cost
		}
	}
	fmt.Fprintln(w, lowest)
}

func CostReal(in, v int) int {
	diff := int(math.Abs(float64(in) - float64(v)))
	var cost int
	for i := 1; i <= diff; i++ {
		cost += i
	}
	return cost
}
