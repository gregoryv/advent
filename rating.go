package advent

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func NewRating(width int) *Rating {
	return &Rating{
		width: width,
		rad:   NewRadiation(width),
		nb:    make(NBits, 0),
	}
}

type Rating struct {
	width int
	rad   *Radiation

	nb NBits

	debug bool
}

func (me *Rating) Parse(r io.Reader) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		me.Write(s.Bytes())
	}
}

func (me *Rating) Write(p []byte) (int, error) {
	me.rad.Write(p)
	me.nb.Write(p)
	return len(p), nil
}

func (me *Rating) Dump() string {
	return me.nb.Dump(me.width)
}

func (me *Rating) Oxygen() Bits {
	b := me.filter(me.nb, me.nb, me.width, 0)
	if me.debug {
		fmt.Println(b.Dump(me.width), "oxygen")
	}
	return b
}

func (me *Rating) filter(last, in NBits, width, i int) Bits {
	if i == width {
		if len(in) == 1 {
			return in[0]
		}
		return last[0] // todo perhaps
	}

	if me.debug {
		fmt.Println(in.Dump(width))
	}

	//
	rad := NewRadiation(width)
	rad.Load(in)
	rest := Keep(in, me.oxygenRating(rad, width, i))

	return me.filter(in, rest, width, i+1)

}

func (me *Rating) oxygenRating(rad *Radiation, width, i int) func(Bits) bool {
	if me.debug {
		k := 1
		if rad.one[i] < rad.zero[i] {
			k = 0
		}
		fmt.Printf("%s%v\n", strings.Repeat(" ", i), k)
	}
	var flag Bits = 1 << (width - i - 1)
	return func(b Bits) bool {
		if rad.one[i] >= rad.zero[i] && Has(b, flag) {
			return true
		}
		if rad.one[i] < rad.zero[i] && !Has(b, flag) {
			return true
		}
		return false
	}
}
