package advent

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Expects binary numbers, one on each line
func ParseNBits(r io.Reader) NBits {
	s := bufio.NewScanner(r)
	nb := make(NBits, 0)
	for s.Scan() {
		nb.Write(s.Bytes())
	}
	return nb
}

type NBits []Bits

func (me *NBits) Write(p []byte) (int, error) {
	b := ParseBitsBytes(p)
	*me = append(*me, b)
	return len(p), nil
}

func ParseBitsBytes(p []byte) Bits {
	return ParseBits(string(p))
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
