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
		cost := alignTo(positions, pos)
		if lowest == 0 || cost < lowest {
			lowest = cost
		}
	}
	fmt.Fprintln(w, lowest)
}

// alignTo returns cost to align input to the given value
func alignTo(input []int, v int) int {
	var sum int
	for _, in := range input {
		sum += int(math.Abs(float64(in) - float64(v)))
	}
	return sum
}
