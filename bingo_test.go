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
	for _, move := range moves {
		v, _ := strconv.Atoi(move)
		g.moves = append(g.moves, v)
	}

	for s.Scan() {
		board := NewBoard()
		for i := 0; i < 5; i++ {
			row := make([]int, 5)
			s.Scan()
			cells := strings.Fields(s.Text())
			for i, str := range cells {
				v, _ := strconv.Atoi(str)
				row[i] = v
			}
			board.values = append(board.values, row)
		}
		g.boards = append(g.boards, board)
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

func NewBoard() *Board {
	return &Board{
		values: make([][]int, 0),
	}
}

type Board struct {
	values [][]int
}

func (me Board) WriteTo(w io.Writer) (int64, error) {
	p, err := nexus.NewPrinter(w)
	for _, row := range me.values {
		for _, cell := range row {
			p.Printf("%2v ", cell)
		}
		p.Println()
	}
	return p.Written, *err
}
