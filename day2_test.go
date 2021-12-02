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
	pos := Navigate(bytes.NewReader(movements))
	fmt.Print(pos.h * pos.d)
	// output: 150
}

func Example_Navigate_input() {
	movements, _ := ioutil.ReadFile("testdata/2.input")
	pos := Navigate(bytes.NewReader(movements))
	fmt.Print(pos.h * pos.d)
	// output: 1580000
}

func Navigate(r io.Reader) *Position {
	s := bufio.NewScanner(r)
	pos := &Position{}
	for s.Scan() {
		line := s.Text()
		if line == "" {
			continue
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

	return pos
}

type Position struct {
	h int
	d int
}
