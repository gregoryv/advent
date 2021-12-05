package advent

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CountIntersections(filename string) {
	fh, err := os.Open(filename)
	shouldNot(err)
	defer fh.Close()
	CountIntersectionsTo(os.Stdout, fh)
}

func CountIntersectionsTo(w io.Writer, r io.Reader) {
	fmt.Fprintln(w, "todo")
}

// ParseLines parses new line separated
//
//   x1,y1 -> x2,y2
//   x1,y1 -> x2,y2
//
func ParseLines(r io.Reader) []Line {
	s := bufio.NewScanner(r)
	lines := make([]Line, 0)
	for s.Scan() {
		l := ParseLine(s.Text())
		lines = append(lines, l)
	}
	return lines
}

// ParseLine parses x1,y1 -> x2,y2
func ParseLine(v string) Line {
	var l Line
	fmt.Scanf("%v,%v -> %v,%v", &l.from.x, &l.from.y, &l.to.x, &l.to.y)
	return l
}

type Grid struct {
	grid [][]int
}

// Intersect returns true if lines 11 and l2 intersect
func (me *Grid) IntersectCount() int {
	var count int
	for y := 0; y < len(me.grid); y++ {
		for x := 0; x < len(me.grid); x++ {
			if me.grid[y][x] > 1 {
				count++
			}
		}
	}
	return count
}

func (me *Grid) Set(l Line) {
	// todo update count in each cell the line crosses
}

type Line struct {
	from, to Pos
}

func (me Line) String() string {
	return fmt.Sprintf("%s -> %s", me.from, me.to)
}

type Pos struct {
	x, y int
}

func (me Pos) String() string {
	return fmt.Sprintf("%v,%v", me.x, me.y)
}
