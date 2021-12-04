package advent

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func Example_PowerCons() {
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