package advent

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestBits(t *testing.T) {
	const (
		F0 Bits = 1 << iota
		F1
		F2
	)
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

func ParseBits(s string) Bits {
	v, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	return Bits(v)
}

type Bits int64

func Set(b, flag Bits) Bits    { return b | flag }
func Clear(b, flag Bits) Bits  { return b &^ flag }
func Toggle(b, flag Bits) Bits { return b ^ flag }
func Has(b, flag Bits) bool    { return b&flag != 0 }
