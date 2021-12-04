package advent

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gregoryv/nexus"
)

// Winner writes score of first winning board to stdout
func Winner(filename string) {
	fh, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer fh.Close()
	g := ParseGame(fh)
	for g.PlayNextMove() {
		board := g.Winner()
		if board != nil {
			break
		}
	}
	fmt.Fprintln(os.Stdout, g.Score())
}

// Looser writes score of last winning board to stdout
func Looser(filename string) {
	fh, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer fh.Close()
	g := ParseGame(fh)

	var lastwin *Board
	var lastnum int
	for g.PlayNextMove() {
		board := g.Winner()
		if board != nil {
			g.RemoveWinners()
			lastwin = board
			lastnum = g.lastNum
		}
	}

	got := lastnum * lastwin.SumUnchecked()
	fmt.Fprintln(os.Stdout, got)
}

func ParseGame(r io.Reader) *Game {
	g := NewGame()
	g.boards = make([]*Board, 0) // reset
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
	g.boards = append(g.boards, b)
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
			i := c*rows + r
			row = Set(row, 1<<(width-i-1))
		}
		g.winflags = append(g.winflags, row)
	}

	// diagonal don't count

	return g
}

type Game struct {
	moves   []int
	move    int // current move
	lastNum int

	boards []*Board

	winflags []Bits // horizontal, vertical
}

func (me *Game) Dump() string {
	var buf bytes.Buffer
	p, _ := nexus.NewPrinter(&buf)
	for i, move := range me.moves {
		if i > 0 {
			p.Print(",")
		}
		p.Printf("%v", move)
	}
	p.Print("\n\n")
	for _, board := range me.boards {
		if me.HasWon(board) {
			p.Println("WINNER")
		}
		board.WriteTo(&buf)
		p.Println()
	}

	for _, row := range me.winflags {
		p.Println(row.Dump(me.boards[0].width))
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
	me.lastNum = num
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
		if me.HasWon(board) {
			return board
		}
	}
	return nil
}

func (me *Game) RemoveWinners() {
	for i := 0; i < len(me.boards); i++ {
		if me.HasWon(me.boards[i]) {
			me.boards = append(me.boards[:i], me.boards[i+1:]...)
		}
	}
}

func (me *Game) Score() int {
	board := me.Winner()
	if board == nil {
		return -1
	}
	log.Println(me.lastNum, board.SumUnchecked())
	return me.lastNum * board.SumUnchecked()
}

func (me *Game) HasWon(b *Board) bool {
	for _, row := range me.winflags {
		if b.Match(row) {
			return true
		}
	}
	return false
}
