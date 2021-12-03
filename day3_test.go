package advent

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

func Example_Radiation() {
	Radiation("testdata/3.given", 5)
	Radiation("testdata/3.input", 12)

	// output:
	// 198
	// 3009600
}

func Radiation(filename string, width int) {
	data := loadLines(filename)

	gamma := make([]byte, width)
	epsilon := make([]byte, width)

	for w := 0; w < width; w++ {
		gamma[w] = '0'
		epsilon[w] = '0'
		var one, zero int

		for _, line := range data {
			if len(line) == 0 { // skip empty
				continue
			}
			if line[w] == '1' {
				one++
			} else {
				zero++
			}
		}
		if one > zero {
			gamma[w] = '1'
		} else {
			epsilon[w] = '1'
		}
	}

	g, _ := strconv.ParseInt(string(gamma), 2, 64)
	e, _ := strconv.ParseInt(string(epsilon), 2, 64)
	fmt.Println(g * e)
}

func loadLines(filename string) [][]byte {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return bytes.Split(data, []byte("\n"))
}
