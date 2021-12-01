package aoc2021

import (
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

func Test_Day1(t *testing.T) {
	cases := []struct {
		in          []int
		exp, expWin int
		msg         string
	}{
		{
			in:     []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
			exp:    7,
			expWin: 5,
			msg:    "given example",
		},
		{
			in:     loadInts(t, "testdata/1.input"),
			exp:    1288,
			expWin: 1311,
			msg:    "my input",
		},
	}
	for _, c := range cases {
		t.Run(c.msg, func(t *testing.T) {
			t.Run("IncreasingInts", func(t *testing.T) {
				got := IncreasingInts(c.in)
				if got != c.exp {
					t.Error("got", got, "exp", c.exp)
				}
			})
			t.Run("IncreasingWindow", func(t *testing.T) {
				got := IncreasingWindow(c.in, 3)
				if got != c.expWin {
					t.Error("got", got, "exp", c.expWin)
				}
			})
		})
	}
}

func IncreasingInts(in []int) int {
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

	// skip first value
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
func loadInts(t *testing.T, filename string) []int {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	lines := strings.Split(string(data), "\n")
	in := make([]int, len(lines))
	for i, line := range lines {
		if line == "" { // skip empty lines
			continue
		}
		in[i], err = strconv.Atoi(line)
		if err != nil {
			t.Fatal(err)
		}
	}
	return in
}
