package advent

import (
	"io"

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
	// init winning rows

	// horizontal
	for r := 0; r < rows; r++ {
		var row Bits
		for c := 0; c < cols; c++ {
			i := r*rows + c
			row = Set(row, 1<<(width-i-1))
		}
		b.winflags = append(b.winflags, row)
	}

	// vertical
	for r := 0; r < rows; r++ {
		var row Bits
		for c := 0; c < cols; c++ {
			i := r * rows
			row = Set(row, 1<<(width-i-1))
		}
		b.winflags = append(b.winflags, row)
	}

	// diagonal don't count
	return b
}

type Board struct {
	values  []int
	checked Bits
	width   int // ie. rows*cols
	rows    int
	cols    int

	winflags []Bits // horizontal, vertical
}

func (me *Board) HasWon() bool {
	for _, flag := range me.winflags {
		if Has(me.checked, flag) {
			return true
		}
	}
	return false
}

func (me *Board) Check(v int) {
	for i, p := range me.values {
		if p == v {
			var flag Bits = 1 << (me.width - i - 1)
			me.checked = Set(me.checked, flag)
		}
	}
}

func (me *Board) IsChecked(i int) bool {
	var flag Bits = 1 << (me.width - i - 1)
	return Has(me.checked, flag)
}

func (me *Board) WriteTo(w io.Writer) (int64, error) {
	p, err := nexus.NewPrinter(w)
	for row := 0; row < me.rows; row++ {
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
