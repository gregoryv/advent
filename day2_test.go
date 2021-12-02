package advent

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

func Example_Navigate_given() {
	movements, _ := ioutil.ReadFile("testdata/2.given")
	pos := Navigate(bytes.NewReader(movements), WithoutAim)
	fmt.Print(pos.h * pos.d)
	// output: 150
}

func Example_Navigate_input() {
	movements, _ := ioutil.ReadFile("testdata/2.input")
	pos := Navigate(bytes.NewReader(movements), WithoutAim)
	fmt.Print(pos.h * pos.d)
	// output: 1580000
}

func Navigate(r io.Reader, calc func(*Position, string)) *Position {
	s := bufio.NewScanner(r)
	pos := &Position{}
	for s.Scan() {
		calc(pos, s.Text())
	}

	return pos
}

func WithoutAim(pos *Position, line string) {
	if line == "" {
		return
	}
	mov := strings.Split(line, " ")
	v, _ := strconv.Atoi(mov[1])
	switch mov[0] {
	case "forward":
		pos.h += v
	case "down":
		pos.d += v
	case "up":
		pos.d -= v
	}
}

type Position struct {
	aim int
	h   int
	d   int
}
