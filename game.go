package advent

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

func ParseGame(r io.Reader) *Game {
	g := NewGame()
	s := bufio.NewScanner(r)
	s.Scan()

	moves := strings.Split(s.Text(), ",")

	for s.Scan() {
		b := NewBoard()
		for r := 0; r < b.rows; r++ {
			s.Scan()
			cells := strings.Fields(s.Text())
			for c := 0; c < b.cols; c++ {
				i := r*b.rows + c
				v, _ := strconv.Atoi(cells[c])
				b.values[i] = v
			}
		}
		g.boards = append(g.boards, b)
	}

	for _, move := range moves {
		v, _ := strconv.Atoi(move)
		g.moves = append(g.moves, v)
	}

	return g
}

func NewGame() *Game {
	g := &Game{
		moves:  make([]int, 0),
		boards: make([]*Board, 0),
	}

	var (
		b     = NewBoard()
		rows  = b.rows
		cols  = b.cols
		width = b.width
	)
	// horizontal
	for r := 0; r < rows; r++ {
		var row Bits
		for c := 0; c < cols; c++ {
			i := r*rows + c
			row = Set(row, 1<<(width-i-1))
		}
		g.winflags = append(g.winflags, row)
	}

	// vertical
	for r := 0; r < rows; r++ {
		var row Bits
		for c := 0; c < cols; c++ {
			i := r * rows
			row = Set(row, 1<<(width-i-1))
		}
		g.winflags = append(g.winflags, row)
	}

	// diagonal don't count

	return g
}

type Game struct {
	moves []int
	move  int // current move

	boards []*Board

	winflags []Bits // horizontal, vertical
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

func (me *Game) PlayNextMove() bool {

	num := me.moves[me.move]
	log.Println("num:", num)
	me.Check(num)
	if me.move < len(me.moves)-1 {
		// there are more moves
		me.move += 1
		return true
	}
	return false
}

func (me *Game) Winner() *Board {
	for _, board := range me.boards {
		for _, row := range me.winflags {
			if board.Match(row) {
				return board
			}
		}
	}
	return nil
}
