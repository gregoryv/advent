package advent

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func Example_PowerCons_given() {
	PowerCons("testdata/3.given", 5) // output: 198
}

func Example_PowerCons_input() {
	PowerCons("testdata/3.input", 12) // output: 3009600
}

func PowerCons(filename string, width int) {
	data, _ := ioutil.ReadFile(filename)
	PowerConsTo(os.Stdout, bytes.NewReader(data), width)
}

func PowerConsTo(w io.Writer, r io.Reader, width int) {
	rad := NewRadiation(width)
	rad.Parse(r)
	fmt.Println(rad.Gamma() * rad.Epsilon())
}

// ----------------------------------------
// part 2

func Example_LifeSupport_given() {
	LifeSupportRating("testdata/3.given", 5) // output: 230
}

func Example_LifeSupport_input() {
	LifeSupportRating("testdata/3.input", 12) // output: 6940518
}

func LifeSupportRating(filename string, width int) {
	data, _ := ioutil.ReadFile(filename)
	LifeSupportRatingTo(os.Stdout, bytes.NewReader(data), width)
}

func LifeSupportRatingTo(w io.Writer, r io.Reader, width int) {
	rating := NewRating(width)
	rating.Parse(r)
	//	rating.debug = true
	fmt.Fprint(w, rating.Oxygen()*rating.CO2Scrub())
}
