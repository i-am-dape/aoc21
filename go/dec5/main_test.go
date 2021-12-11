package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/i-am-dape/aoc21/go/test"
)

type Point struct {
	x, y int
}

type Segment struct {
	beginning, end Point
}

func abs(n int) int {
	if n > 0 {
		return n
	}

	return -1 * n
}

func delta(v1, v2 int) (int, int) {
	d := v1 - v2
	if d < 0 {
		return v2 - v1, -1
	}
	return d, 1
}

func (s *Segment) DeltaX() (int, int) {
	return delta(s.end.x, s.beginning.x)
}

func (s *Segment) DeltaY() (int, int) {
	return delta(s.end.y, s.beginning.y)
}

func (s *Segment) IsHorizontal() bool {
	return s.beginning.y == s.end.y
}

func (s *Segment) IsVertical() bool {
	return s.beginning.x == s.end.x
}

func (s *Segment) IsDiagonal() bool {
	dx := s.end.x - s.beginning.x
	if dx < 0 {
		dx = s.beginning.x - s.end.x
	}

	dy := s.end.y - s.beginning.y
	if dy < 0 {
		dy = s.beginning.y - s.end.y
	}

	return dx == dy
}

func TestSegmentDelta(t *testing.T) {
	tests := []struct {
		s     Segment
		dx    int
		dxDir int
		dy    int
		dyDir int
	}{
		{
			Segment{Point{}, Point{}},
			0,
			1,
			0,
			1,
		},
		{
			Segment{Point{0, 0}, Point{5, 5}},
			5,
			1,
			5,
			1,
		},
		{
			Segment{Point{5, 5}, Point{0, 0}},
			5,
			-1,
			5,
			-1,
		},
		{
			Segment{Point{0, 0}, Point{0, 9}},
			0,
			1,
			9,
			1,
		},
		{
			Segment{Point{9, 4}, Point{3, 4}},
			6,
			-1,
			0,
			1,
		},
	}

	for _, tc := range tests {
		t.Log(tc)

		dx, dxDir := tc.s.DeltaX()
		dy, dyDir := tc.s.DeltaY()

		if dx != tc.dx || dxDir != tc.dxDir || dy != tc.dy || dyDir != tc.dyDir {
			t.Error(dx, dxDir, dy, dyDir)
		}
	}
}

func ParseSegment(txt string) (Segment, error) {
	outer := strings.Split(txt, "->")
	beginningTxt := strings.Split(outer[0], ",")
	endTxt := strings.Split(outer[1], ",")

	bx, err := strconv.ParseInt(strings.TrimSpace(beginningTxt[0]), 10, 32)
	if err != nil {
		return Segment{}, err
	}

	by, err := strconv.ParseInt(strings.TrimSpace(beginningTxt[1]), 10, 32)
	if err != nil {
		return Segment{}, err
	}

	ex, err := strconv.ParseInt(strings.TrimSpace(endTxt[0]), 10, 32)
	if err != nil {
		return Segment{}, err
	}

	ey, err := strconv.ParseInt(strings.TrimSpace(endTxt[1]), 10, 32)
	if err != nil {
		return Segment{}, err
	}

	return Segment{Point{int(bx), int(by)}, Point{int(ex), int(ey)}}, nil
}

type Grid struct {
	table [][]int
}

func (g *Grid) String() string {
	b := &strings.Builder{}
	fmt.Fprintf(b, "%dx%d", len(g.table), len(g.table))
	fmt.Fprintln(b)
	fmt.Fprintln(b)

	for _, y := range g.table {
		fmt.Fprintln(b, y)
	}

	return b.String()
}

func (g *Grid) mark(x, y int) int {
	cur := g.table[y][x]
	cur++
	g.table[y][x] = cur
	return cur
}

func (g *Grid) Mark(s Segment) {
	fmt.Printf("Segment: %v ", s)

	switch {
	case s.IsHorizontal():
		dx, dir := s.DeltaX()
		fmt.Print("H dx:", dx, " dir:", dir, " |")
		for i := 0; i <= dx; i++ {
			x := s.beginning.x + (i * dir)
			fmt.Print(" ", x)
			g.mark(x, s.beginning.y)
		}
		fmt.Println()

	case s.IsVertical():
		dy, dir := s.DeltaY()
		fmt.Print("V dy:", dy, " dir:", dir, " |")
		for i := 0; i <= dy; i++ {
			y := s.beginning.y + (i * dir)
			fmt.Print(" ", y)
			g.mark(s.beginning.x, y)
		}
		fmt.Println()

	case s.IsDiagonal():
		dx, dxdir := s.DeltaX()
		fmt.Print("DH dx:", dx, " dir:", dxdir, " |")
		dy, dydir := s.DeltaY()
		fmt.Print("DV dy:", dy, " dir:", dydir, " |")
		if dy != dx {
			panic("dx and dy must be identical for diagonal lines")
		}

		for i := 0; i <= dx; i++ {
			x := s.beginning.x + (i * dxdir)
			y := s.beginning.y + (i * dydir)
			fmt.Print(" (", x, ",", y, ")")
			g.mark(x, y)
		}
		fmt.Println()

	default:
		panic("only horizontal, vertical and diagonal lines supported")
	}
}

func NewGrid(n int) *Grid {
	rows := make([][]int, n)
	for r := 0; r < n; r++ {
		rows[r] = make([]int, n)
	}
	return &Grid{rows}
}

func MaxXY(segments []Segment) int {
	max := func(v1, v2 int) int {
		if v1 > v2 {
			return v1
		}

		return v2
	}

	x := 0
	y := 0

	for _, segment := range segments {
		x = max(max(segment.beginning.x, segment.end.x), x)
		y = max(max(segment.beginning.y, segment.end.y), y)
	}

	return max(x, y)
}

func TestDec5(t *testing.T) {
	test.Run(func() {
		segements, err := test.Read(ParseSegment)
		if err != nil {
			t.Fatal(err)
		}
		g := NewGrid(MaxXY(segements) + 1)

		for _, s := range segements {
			if s.IsHorizontal() || s.IsVertical() {
				g.Mark(s)
			}
		}

		overlaps := 0
		for y := 0; y < len(g.table); y++ {
			for x := 0; x < len(g.table[y]); x++ {
				if g.table[y][x] > 1 {
					overlaps++
				}
			}
		}

		t.Log(g)
		test.Check(t, overlaps, 5, 5147)
	})
}

func TestIsDiagonal(t *testing.T) {
	tests := []struct {
		s    Segment
		want bool
	}{
		{
			Segment{Point{1, 1}, Point{3, 3}},
			true,
		},
		{
			Segment{Point{10, 10}, Point{3, 3}},
			true,
		},
		{
			Segment{Point{1, 10}, Point{3, 3}},
			false,
		},
		{
			Segment{Point{2, 0}, Point{6, 4}},
			true,
		},
		{
			Segment{Point{8, 0}, Point{0, 8}},
			true,
		},
	}

	for _, tc := range tests {
		t.Log(tc)
		got := tc.s.IsDiagonal()
		if got != tc.want {
			t.Fatal(tc.s, got, tc.want)
		}
	}
}

func TestDec5_2(t *testing.T) {
	test.Run(func() {
		segements, err := test.Read(ParseSegment)
		if err != nil {
			t.Fatal(err)
		}
		g := NewGrid(MaxXY(segements) + 1)
		t.Log(g)

		for _, s := range segements {
			if s.IsHorizontal() || s.IsVertical() || s.IsDiagonal() {
				g.Mark(s)
			} else {
				fmt.Println("skip", s)
			}
		}

		overlaps := 0
		for y := 0; y < len(g.table); y++ {
			for x := 0; x < len(g.table[y]); x++ {
				if g.table[y][x] > 1 {
					overlaps++
				}
			}
		}

		t.Log(g)
		test.Check(t, overlaps, 12, 16925)
	})
}

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}
