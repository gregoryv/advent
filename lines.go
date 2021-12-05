package advent

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/gregoryv/nexus"
)

func CountHVIntersections(filename string, x, y int) {
	CountIntersections(filename, x, y, func(l Line) bool {
		return l.IsHorizontal() || l.IsVertical()
	})
}

func CountIntersections(filename string, x, y int, keep func(Line) bool) {
	fh, err := os.Open(filename)
	shouldNot(err)
	defer fh.Close()
	CountIntersectionsTo(os.Stdout, fh, x, y, keep)
}

func CountIntersectionsTo(w io.Writer, r io.Reader, x, y int, keep func(Line) bool) {
	g := NewGrid(x, y)
	lines := ParseLines(r)
	for _, line := range lines {
		if keep(line) {
			g.Set(line)
		}
	}
	fmt.Fprintln(w, g.IntersectCount())
	if debugOn {
		debug.Logf("\n%s", g.Dump())
	}
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
	v = strings.ReplaceAll(v, " -> ", " ")
	v = strings.ReplaceAll(v, ",", " ")
	var l Line
	n, _ := fmt.Sscan(v, &l.from.x, &l.from.y, &l.to.x, &l.to.y)
	debug.Log("ParseLine:", n, v)
	return l
}

func NewGrid(x, y int) *Grid {
	g := &Grid{
		grid: make([][]int, 0, y),
	}
	for i := 0; i < y; i++ {
		g.grid = append(g.grid, make([]int, x))
	}
	return g
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
	l.Walk(me.SetPos)
}

func (me *Grid) SetPos(p Pos) {
	me.grid[p.y][p.x]++
}

func (me *Grid) Dump() string {
	var buf bytes.Buffer
	p, _ := nexus.NewPrinter(&buf)
	for y := 0; y < len(me.grid); y++ {
		for x := 0; x < len(me.grid); x++ {
			p.Print(me.grid[y][x])
		}
		p.Println()
	}
	return buf.String()
}

// ----------------------------------------

type Line struct {
	from, to Pos
}

func (me Line) String() string {
	return fmt.Sprintf("%s -> %s", me.from, me.to)
}

// Walk calls fn with each position from to start
func (l Line) Walk(fn func(p Pos)) {
	fn(l.from)
	prev := l.from
	for {
		next := prev.Next(l.to)
		if prev.x == next.x && prev.y == next.y {
			break
		}
		fn(next)
		prev = next
	}
}

func (l Line) IsVertical() bool   { return l.from.y == l.to.y }
func (l Line) IsHorizontal() bool { return l.from.x == l.to.x }

type Pos struct {
	x, y int
}

func (me Pos) String() string {
	return fmt.Sprintf("%d,%d", me.x, me.y)
}

func (me Pos) Next(p Pos) Pos {
	next := me // copy

	// direction of lines
	switch {
	case p.x > me.x:
		next.x++
	case p.x < me.x:
		next.x--
	}

	switch {
	case p.y > me.y:
		next.y++
	case p.y < me.y:
		next.y--
	}
	return next
}
