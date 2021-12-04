package advent

import (
	"bufio"
	"io"
)

func NewRadiation(width int) *Radiation {
	return &Radiation{
		width: width,

		one:  make([]int, width),
		zero: make([]int, width),
	}
}

type Radiation struct {
	width int

	one  []int
	zero []int

	gamma   Bits
	epsilon Bits
}

func (me *Radiation) Parse(r io.Reader) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		me.Write(s.Bytes())
	}
}

func (me *Radiation) Load(data [][]byte) {
	for _, line := range data {
		me.Write(line)
	}
}

// p must be exact width
func (me *Radiation) Write(p []byte) (int, error) {
	if len(p) == 0 { // skip empty
		return 0, nil
	}
	b := ParseBitsBytes(p)

	for i := 0; i < me.width; i++ {
		var flag Bits = 1 << (me.width - i - 1)
		if Has(b, flag) {
			me.one[i]++
		} else {
			me.zero[i]++
		}
	}
	return len(p), nil
}

func (me *Radiation) Gamma() int64 {
	me.update()
	return int64(me.gamma)
}

func (me *Radiation) Epsilon() int64 {
	me.update()
	return int64(me.epsilon)
}

func (me *Radiation) update() {
	for i := 0; i < me.width; i++ {
		var flag Bits = 1 << (me.width - i - 1)
		if me.one[i] > me.zero[i] {
			me.gamma = Set(me.gamma, flag)
		} else {
			me.epsilon = Set(me.epsilon, flag)
		}
	}
}
