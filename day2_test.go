package advent

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func Example_Navigate() {
	Navigate("testdata/2.given", WithoutAim)
	Navigate("testdata/2.input", WithoutAim)
	Navigate("testdata/2.given", WithAim)
	Navigate("testdata/2.input", WithAim)
	// output:
	// 150
	// 1580000
	// 900
	// 1251263225
}

func Navigate(filename string, calc func(*Position, string)) {
	movements, _ := ioutil.ReadFile(filename)
	NavigateTo(os.Stdout, bytes.NewReader(movements), calc)
}

func NavigateTo(w io.Writer, r io.Reader, calc func(*Position, string)) {
	s := bufio.NewScanner(r)
	pos := &Position{}
	for s.Scan() {
		calc(pos, s.Text())
	}
	fmt.Fprintln(w, pos)
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

func WithAim(pos *Position, line string) {
	if line == "" {
		return
	}
	mov := strings.Split(line, " ")
	v, _ := strconv.Atoi(mov[1])
	switch mov[0] {
	case "forward":
		pos.h += v
		pos.d = pos.d + (v * pos.aim)
	case "down":
		pos.aim += v
	case "up":
		pos.aim -= v

	}
}

type Position struct {
	aim int
	h   int
	d   int
}

func (me *Position) String() string {
	return fmt.Sprintf("%v", me.h*me.d)
}
