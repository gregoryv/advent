package advent

import "testing"

func TestParseLine(t *testing.T) {
	SetDebug(t)
	defer SetDebug(nil)
	in := "1,1 -> 3,3"
	l := ParseLine(in)

	got := l.String()
	exp := in
	if got != exp {
		t.Errorf("got %v, expected %v", got, exp)
	}
}
