package advent

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func PowerCons(filename string, width int) {
	data, _ := ioutil.ReadFile(filename)
	PowerConsTo(os.Stdout, bytes.NewReader(data), width)
}

func PowerConsTo(w io.Writer, r io.Reader, width int) {
	rad := NewRadiation(width)
	rad.Parse(r)
	fmt.Println(rad.Gamma() * rad.Epsilon())
}

func LifeSupportRating(filename string, width int) {
	data, _ := ioutil.ReadFile(filename)
	LifeSupportRatingTo(os.Stdout, bytes.NewReader(data), width)
}

func LifeSupportRatingTo(w io.Writer, r io.Reader, width int) {
	rating := NewRating(width)
	rating.Parse(r)
	fmt.Fprintln(w, rating.Oxygen()*rating.CO2Scrub())
}

// ----------------------------------------

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
	b := me.filter(me.oxygenRating, me.nb, me.nb, me.width, 0)
	if debugOn {
		fmt.Println(b.Dump(me.width), "oxygen")
	}
	return b
}

func (me *Rating) oxygenRating(rad *Radiation, width, i int) func(Bits) bool {
	if debugOn {
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

func (me *Rating) CO2Scrub() Bits {
	b := me.filter(me.co2scrub, me.nb, me.nb, me.width, 0)
	if debugOn {
		fmt.Println(b.Dump(me.width), "co2")
	}
	return b
}

func (me *Rating) co2scrub(rad *Radiation, width, i int) func(Bits) bool {
	if debugOn {
		k := 0
		if rad.one[i] < rad.zero[i] {
			k = 1
		}
		fmt.Printf("%s%v\n", strings.Repeat(" ", i), k)
	}
	var flag Bits = 1 << (width - i - 1)
	return func(b Bits) bool {
		if rad.one[i] < rad.zero[i] && Has(b, flag) {
			return true
		}
		if rad.one[i] >= rad.zero[i] && !Has(b, flag) {
			return true
		}
		return false
	}
}

func (me *Rating) filter(match matchFunc, last, in NBits, width, i int) Bits {
	if len(in) == 1 {
		return in[0]
	}
	if len(last) == 0 {
		return 0
	}
	if i == width {
		log.Println("stopped")
		return 0
	}
	if debugOn {
		fmt.Println(in.Dump(width))
	}
	//
	rad := NewRadiation(width)
	rad.Load(in)
	rest := Keep(in, match(rad, width, i))

	return me.filter(match, in, rest, width, i+1)
}

type matchFunc func(rad *Radiation, width, i int) func(Bits) bool
