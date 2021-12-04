package advent

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/gregoryv/nexus"
)

func Test(t *testing.T) {
	r, _ := os.Open("testdata/4.given")
	g := ParseGame(r)

	t.Error(g.Dump())
}

func ParseGame(r io.Reader) *Game {
	g := NewGame()
	s := bufio.NewScanner(r)
	s.Scan()

	moves := strings.Split(s.Text(), ",")

	for s.Scan() {
		board := NewBoard()
		for i := 0; i < 5; i++ {
			s.Scan()
			cells := strings.Fields(s.Text())
			for _, str := range cells {
				v, _ := strconv.Atoi(str)
				board.values = append(board.values, v)
			}
		}
		g.boards = append(g.boards, board)
	}

	for _, move := range moves[:5] {
		v, _ := strconv.Atoi(move)
		g.moves = append(g.moves, v)
		g.Check(v)
	}

	return g
}

func NewGame() *Game {
	return &Game{
		moves:  make([]int, 0),
		boards: make([]*Board, 0),
	}
}

type Game struct {
	moves  []int
	boards []*Board
}

func (me *Game) Dump() string {
	var buf bytes.Buffer
	for i, move := range me.moves {
		if i > 0 {
			buf.WriteString(",")
		}
		buf.WriteString(fmt.Sprintf("%v", move))
	}
	buf.WriteString("\n\n")
	for _, board := range me.boards {
		board.WriteTo(&buf)
		buf.WriteString("\n")
	}
	return buf.String()
}

func (me *Game) Check(v int) {
	for _, board := range me.boards {
		board.Check(v)
	}
}

func NewBoard() *Board {
	return &Board{
		values: make([]int, 0),
		width:  5 * 5,
		rows:   5,
		cols:   5,
	}
}

type Board struct {
	values []int

	checked Bits

	width, rows, cols int
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
