package aoc2021

import (
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

func Test_Dai1_part1(t *testing.T) {
	load := func(filename string) []int {
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
	cases := []struct {
		in  []int
		exp int
		msg string
	}{
		{
			in:  []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
			exp: 7,
			msg: "given example",
		},
		{
			in:  load("testdata/1.input"),
			exp: 1288,
			msg: "my input",
		},
	}
	for _, c := range cases {
		t.Run(c.msg, func(t *testing.T) {
			got := Day1_part1(c.in)
			if got != c.exp {
				t.Error("got", got, "exp", c.exp)
			}
		})
	}
}
