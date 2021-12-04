package advent

import (
	"fmt"
	"strings"
	"testing"
)

func TestNBits(t *testing.T) {
	t.Run("ParseNBits", func(t *testing.T) {
		r := strings.NewReader("101\n010\n001") // without trailing newline
		nb := ParseNBits(r)
		if len(nb) != 3 {
			t.Error("len:", len(nb))
		}
		if nb[1] != F1 {
			t.Errorf("%03b", nb[1])
		}
	})

	t.Run("Write", func(t *testing.T) {
		nb := make(NBits, 0)
		nb.Write([]byte("101"))
		nb.Write([]byte("010"))
		nb.Write([]byte("001"))

		if len(nb) != 3 {
			t.Fail()
		}
		if nb[1] != F1 {
			t.Errorf("%03b", nb[1])
		}
	})
}

// ----------------------------------------

func TestBits(t *testing.T) {
	t.Run("ParseBits", func(t *testing.T) {
		b := ParseBits("0101")
		got := fmt.Sprintf("%03b", b)
		exp := "101"
		if got != exp {
			t.Errorf("got %s, expected %s", got, exp)
		}
	})

	t.Run("Set", func(t *testing.T) {
		b := Set(Set(0, F0), F2)
		got := fmt.Sprintf("%03b", b)
		exp := "101"
		if got != exp {
			t.Errorf("got %s, expected %s", got, exp)
		}
	})
	t.Run("Clear", func(t *testing.T) {
		b := Set(Set(0, F0), F2)
		b = Clear(b, F2)
		got := fmt.Sprintf("%03b", b)
		exp := "001"
		if got != exp {
			t.Errorf("got %s, expected %s", got, exp)
		}
	})
	t.Run("Toggle", func(t *testing.T) {
		b := Set(Set(0, F0), F2)
		b = Toggle(b, F2)
		b = Toggle(b, F1)
		got := fmt.Sprintf("%03b", b)
		exp := "011"
		if got != exp {
			t.Errorf("got %s, expected %s", got, exp)
		}
	})
	t.Run("Set", func(t *testing.T) {
		var b Bits = 0
		if Has(b, F1) {
			t.Fatal()
		}
		b = Set(b, F1)
		got := fmt.Sprintf("%03b", b)
		exp := "010"
		if !Has(b, F1) {
			t.Errorf("got %s, expected %s", got, exp)
		}
	})
}

const (
	F0 Bits = 1 << iota
	F1
	F2
)
