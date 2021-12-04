package advent

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
)

func Keep(nb NBits, match func(Bits) bool) NBits {
	out := make(NBits, 0)
	for _, b := range nb {
		if match(b) {
			out = append(out, b)
		}
	}
	return out
}

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

func (me *NBits) Dump(width int) string {
	var buf bytes.Buffer
	format := fmt.Sprintf("%%0%vb\n", width)
	for _, b := range *me {
		fmt.Fprintf(&buf, format, b)
	}
	return buf.String()
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

func (b Bits) Dump(width int) string {
	format := fmt.Sprintf("%%0%vb", width)
	return fmt.Sprintf(format, b)
}

// SetIndex from right
func SetIndex(b Bits, v ...uint8) Bits {
	for _, v := range v {
		b = Set(b, (1 << v))
	}
	return b
}

func Set(b, flag Bits) Bits    { return b | flag }
func Clear(b, flag Bits) Bits  { return b &^ flag }
func Toggle(b, flag Bits) Bits { return b ^ flag }
func Has(b, flag Bits) bool    { return b&flag != 0 }
func Match(b, flag Bits) bool  { return b&flag == flag }
