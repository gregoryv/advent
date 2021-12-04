package advent

import (
	"testing"
)

func TestBoard_Match(t *testing.T) {
	b := NewBoard()
	b.Check(1, 2, 3, 4)
	row := SetIndex(0, 20, 21, 22, 23, 24)

	if b.Match(row) {
		t.Error(row.Dump(25))
		t.Errorf("%s", b.Dump())
	}

	b.Check(5)
	if !b.Match(row) {
		t.Fail()
	}
}
