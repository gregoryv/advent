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

func xTestCountHVInteractions(t *testing.T) {
	SetDebug(t)
	defer SetDebug(nil)
	CountHVIntersections("testdata/5.given", 10, 10)
	t.Fail()
}

func TestPos_Next(t *testing.T) {
	t.Run("left to right", func(t *testing.T) {
		from, to := Pos{0, 0}, Pos{3, 0}
		got := from.Next(to)
		got = got.Next(to)
		got = got.Next(to)
		if got.x != 3 {
			t.Error(got.String())
		}
	})

	t.Run("right to left", func(t *testing.T) {
		from, to := Pos{3, 0}, Pos{0, 0}
		got := from.Next(to)
		got = got.Next(to)
		got = got.Next(to)
		if got.x != 0 {
			t.Error(got.String())
		}
	})

	t.Run("down", func(t *testing.T) {
		from, to := Pos{0, 0}, Pos{0, 3}
		got := from.Next(to)
		got = got.Next(to)
		got = got.Next(to)
		if got.y != 3 {
			t.Error(got.String())
		}
	})

	t.Run("up", func(t *testing.T) {
		from, to := Pos{0, 3}, Pos{0, 0}
		got := from.Next(to)
		got = got.Next(to)
		got = got.Next(to)
		if got.y != 0 {
			t.Error(got.String())
		}
	})

}
