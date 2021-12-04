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
