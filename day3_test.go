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
}

func LifeSupportRating(filename string, width int) {
	r, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()
	lifeSupportRating(os.Stdout, r, width)
}

func lifeSupportRating(w io.Writer, r io.Reader, width int) {

}

func NewRating() *Rating {
	return &Rating{}
}

type Rating struct {
	oxygen    []byte
	scrubbing []byte
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
