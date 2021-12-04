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
)

func Test(t *testing.T) {
	r, _ := os.Open("testdata/4.given")
	g := ParseGame(r)

	for g.PlayNextMove() {
		board := g.Winner()
		if board != nil {
			t.Fatal(g.Dump())
		}
	}

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

	for _, move := range moves {
		v, _ := strconv.Atoi(move)
		g.moves = append(g.moves, v)
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
	moves []int
	move  int // current move

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

func (me *Game) PlayNextMove() bool {
	me.Check(me.moves[me.move])
	if me.move < len(me.moves)-1 {
		// there are more moves
		me.move += 1
		return true
	}
	return false
}

func (me *Game) Winner() *Board {
	for _, board := range me.boards {
		if board.HasWon() {
			return board
		}
	}
	return nil
}
