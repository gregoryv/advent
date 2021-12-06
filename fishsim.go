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
		lanterns: make([]int, 9),
	}
	ages := toInts(initial)
	for _, v := range ages {
		sim.lanterns[v] += 1
	}
	sim.Run(days)
	fmt.Fprintln(w, sim.Total())
}

type FishSim struct {
	lanterns []int // 0 .. 8
}

func (me *FishSim) Run(days int) {
	l := len(me.lanterns)
	for d := 0; d < days; d++ {
		parents := me.lanterns[0]
		born := parents

		// shift all lanterns, ie. they are one day closer to being parents
		for i := 1; i < l; i++ {
			me.lanterns[i-1] = me.lanterns[i]
		}

		// current parents will be parents again in 6 days
		me.lanterns[6] += parents

		// add new born to the group
		me.lanterns[8] = born
	}
}

func (me *FishSim) Total() int {
	var sum int
	for _, v := range me.lanterns {
		sum += v
	}
	return sum
}
