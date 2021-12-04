package advent

import (
	"bytes"
	"io"
	"strings"

	"github.com/gregoryv/nexus"
)

func NewBoard() *Board {
	rows, cols := 5, 5
	width := rows * cols
	b := &Board{
		values: make([]int, 0),
		width:  width,
		rows:   rows,
		cols:   cols,
	}
	// default to values 1 .. width
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			i := r*rows + c
			b.values = append(b.values, i+1)
		}
	}
	// init winning rows , todo move to game as these are game rules

	return b
}

type Board struct {
	values  []int
	checked Bits
	width   int // ie. rows*cols
	rows    int
	cols    int
}

func (me *Board) Check(v ...int) {
	for _, v := range v {
		for i, p := range me.values {
			if p == v {
				var flag Bits = 1 << (me.width - i - 1)
				me.checked = Set(me.checked, flag)
			}
		}
	}
}

func (me *Board) SumUnchecked() int {
	var sum int
	for i := 0; i < me.width; i++ {
		if !me.IsChecked(i) {
			sum += me.values[i]
		}
	}
	return sum
}

func (me *Board) IsChecked(i int) bool {
	var flag Bits = 1 << (me.width - i - 1)
	return Has(me.checked, flag)
}

func (me *Board) Match(flag Bits) bool {
	return Match(me.checked, flag)
}

func (me *Board) Dump() string {
	var buf bytes.Buffer
	me.WriteTo(&buf)
	return buf.String()
}

func (me *Board) WriteTo(w io.Writer) (int64, error) {
	p, err := nexus.NewPrinter(w)
	indent := strings.Repeat(" ", 20)
	p.Println(me.checked.Dump(me.width))
	for row := 0; row < me.rows; row++ {
		p.Print(indent)
		for cell := 0; cell < me.cols; cell++ {
			i := row*me.rows + cell
			if !me.IsChecked(i) {
				p.Printf("%2v ", me.values[i])
			} else {
				p.Printf("%s%2v%s ", RED, me.values[i], NOCOLOR)
			}
		}
		p.Println()
	}
	return p.Written, *err
}

const (
	RED     = "\033[0;31m"
	NOCOLOR = "\033[0m"
)
