package advent

import (
	"os"
	"testing"
)

func Test(t *testing.T) {
	r, _ := os.Open("testdata/4.given")
	g := ParseGame(r)

	for g.PlayNextMove() {
		board := g.Winner()
		if board != nil {
			break
		}
	}
	got := g.Score()
	exp := 4512
	if got != exp {
		t.Errorf("got %v, expected %v", got, exp)
	}
}
