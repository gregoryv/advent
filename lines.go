package advent

import (
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
	// todo
	return nil
}

// ParseLine parses x1,y1 -> x2,y2
func ParseLine(v string) Line {
	var l Line
	// todo
	return l
}

type Grid struct {
	grid [][]int
}

// Intersect returns true if lines 11 and l2 intersect
func (me *Grid) IntersectCount() bool {
	// todo find intersection, ie. find all cells in a grid with a count > 1
	return false
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
