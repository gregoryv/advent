package advent

import (
	"bufio"
	"fmt"
	"io"
	"log"
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
	b := me.filter(me.nb, me.nb, me.width, me.width)
	if me.debug {
		fmt.Println(b.Dump(me.width), "oxygen")
	}
	return b
}

func (me *Rating) filter(last, in NBits, width, i int) Bits {
	var flag Bits = 1 << (i - 1)
	if me.debug {
		fmt.Println(in.Dump(width))
		fmt.Println(flag.Dump(width), "flag")
		fmt.Println()
	}
	if len(in) == 1 {
		return in[0]
	}
	if i == -1 {
		log.Fatal("stopped")
	}
	//

	rest := Keep(in, oxygenRating(flag))
	return me.filter(in, rest, width, i-1)

}

func oxygenRating(flag Bits) func(Bits) bool {
	return func(b Bits) bool {
		return Has(b, flag)
	}
}
