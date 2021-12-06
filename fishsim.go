package advent

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func LanternFish(filename string, days int) {
	fh, err := os.Open(filename)
	shouldNot(err)
	defer fh.Close()
	LanternFishTo(os.Stdout, fh, days)
}

func LanternFishTo(w io.Writer, r io.Reader, days int) {
	var buf bytes.Buffer
	io.Copy(&buf, r)
	initial := strings.Split(buf.String(), ",")

	sim := &FishSim{
		lanterns: toInts(initial),
	}
	sim.Run(days)
	fmt.Fprintln(w, len(sim.lanterns))
}

type FishSim struct {
	lanterns []int
}

func (me *FishSim) Run(days int) {
	for d := 0; d < days; d++ {
		born := make([]int, 0)
		for i, age := range me.lanterns {
			switch age {
			case 0:
				born = append(born, 8)
				me.lanterns[i] = 6
			default:
				me.lanterns[i]--
			}
		}
		if len(born) > 0 {
			me.lanterns = append(me.lanterns, born...)
		}
	}
}
