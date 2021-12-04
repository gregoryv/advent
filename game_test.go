package advent

import (
	"os"
	"testing"
)

func TestGame_PlayNextMove(t *testing.T) {
	r, _ := os.Open("testdata/4.input")
	g := ParseGame(r)

	for g.PlayNextMove() {
		board := g.Winner()
		if board != nil {
			break
		}
	}
	got := g.Score()
	exp := 16716
	if got != exp {
		t.Errorf("got %v, expected %v", got, exp)
		//		t.Error(g.Dump())
	}
}

func TestGame_Rules(t *testing.T) {
	g := NewGame()

	if false {
		t.Error(g.Dump())
	}
}

func TestGame_LastBoard(t *testing.T) {
	r, _ := os.Open("testdata/4.input")
	g := ParseGame(r)

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
	exp := 4880
	if got != exp {
		t.Errorf("got %v, expected %v", got, exp)
		//		t.Error(g.Dump())
	}
}
