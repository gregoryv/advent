package advent

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func Example_CalcRadiation() {
	CalcRadiation("testdata/3.given", 5)
	CalcRadiation("testdata/3.input", 12)

	LifeSupportRating("testdata/3.given", 5)
	LifeSupportRating("testdata/3.input", 12)

	// output:
	// 198
	// 3009600
	// 230
	// ?
}

// ----------------------------------------

func LifeSupportRating(filename string, width int) {
	data, _ := ioutil.ReadFile(filename)
	LifeSupportRatingTo(os.Stdout, bytes.NewReader(data), width)
}

func LifeSupportRatingTo(w io.Writer, r io.Reader, width int) {
	rating := NewRating(width)
	rating.Parse(r)
	fmt.Fprintln(w, rating.LifeSupport())
}

func NewRating(width int) *Rating {
	return &Rating{
		width: width,
		data:  make([][]byte, 0),
	}
}

type Rating struct {
	width    int
	data     [][]byte
	oxygen   []byte
	co2scrub []byte
}

func (me *Rating) Parse(r io.Reader) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		me.Write(s.Bytes())
	}
	me.oxygen = me.filter(me.data, findOxygen, 0)
	me.co2scrub = me.filter(me.data, findCO2scrub, 0)
}

var last [][]byte

func (me *Rating) filter(in [][]byte, keep keepFunc, i int) []byte {
	if len(in) == 1 {
		return in[0]
	}
	if len(in) == 0 || i == me.width { // probably all last where the same
		return last[0]
	}

	last = in
	k := keep(in, i)
	rest := make([][]byte, 0)
	for _, line := range in {
		if line[i] == k {
			rest = append(rest, line)
		}
	}

	return me.filter(rest, keep, i+1)
}

func findOxygen(in [][]byte, i int) byte {
	width := len(in[0])
	rad := NewRadiation(width)
	rad.Load(in)
	if i >= width {
		log.Fatal(dump(in))
	}
	if rad.one[i] >= rad.zero[i] {
		return '1'
	}
	return '0'

}

func findCO2scrub(in [][]byte, i int) byte {
	width := len(in[0])
	rad := NewRadiation(width)
	rad.Load(in)

	if rad.one[i] < rad.zero[i] {
		return '1'
	}
	return '0'
}

type keepFunc func([][]byte, int) byte

func (me *Rating) Write(p []byte) (int, error) {
	if len(p) == 0 { // skip empty
		return 0, nil
	}
	me.data = append(me.data, p)
	return len(p), nil
}

func (me *Rating) Oxygen() int64 {
	v, _ := strconv.ParseInt(string(me.oxygen), 2, 64)
	return v
}

func (me *Rating) CO2scrub() int64 {
	v, _ := strconv.ParseInt(string(me.co2scrub), 2, 64)
	return v
}

func (me *Rating) LifeSupport() int64 {
	return me.Oxygen() * me.CO2scrub()
}

func dump(in [][]byte) string {
	var buf bytes.Buffer
	for _, word := range in {
		buf.Write(word)
		buf.WriteString("\n")
	}
	return buf.String()
}

// ----------------------------------------

func CalcRadiation(filename string, width int) {
	data, _ := ioutil.ReadFile(filename)
	CalcRadiationTo(os.Stdout, bytes.NewReader(data), width)
}

func CalcRadiationTo(w io.Writer, r io.Reader, width int) {
	rad := NewRadiation(width)
	rad.Parse(r)
	fmt.Println(rad.Gamma() * rad.Epsilon())
}

func NewRadiation(width int) *Radiation {
	return &Radiation{
		one:  make([]int, width),
		zero: make([]int, width),

		gamma:   bytes.Repeat([]byte("0"), width),
		epsilon: bytes.Repeat([]byte("0"), width),
	}
}

type Radiation struct {
	one  []int
	zero []int

	gamma   []byte
	epsilon []byte
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

	for i := 0; i < len(me.gamma); i++ {
		if p[i] == '1' {
			me.one[i]++
		} else {
			me.zero[i]++
		}
	}
	return len(p), nil
}

func (me *Radiation) Update() {
	for i := 0; i < len(me.gamma); i++ {
		if me.one[i] > me.zero[i] {
			me.gamma[i] = '1'
		} else {
			me.epsilon[i] = '1'
		}
	}
}

func (me *Radiation) Gamma() int64 {
	me.Update()
	v, _ := strconv.ParseInt(string(me.gamma), 2, 64)
	return v
}
func (me *Radiation) Epsilon() int64 {
	me.Update()
	v, _ := strconv.ParseInt(string(me.epsilon), 2, 64)
	return v
}

// ----------------------------------------

func loadLines(r io.Reader) [][]byte {
	data, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	return bytes.Split(data, []byte("\n"))
}
