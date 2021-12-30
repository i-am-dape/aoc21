package main

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/i-am-dape/aoc21/go/test"
)

type HM struct {
	// 0,0 -> top left
	data [][]int
}

func (h *HM) Clone() *HM {
	c := &HM{[][]int{}}
	for _, r := range h.data {
		rc := make([]int, len(r))
		copy(rc, r)
		c.data = append(c.data, rc)
	}
	return c
}

type neighbour struct {
	c coord
	v int
}

func (h *HM) Neighbours(c coord) []neighbour {
	left := coord{c.x - 1, c.y}
	right := coord{c.x + 1, c.y}
	above := coord{c.x, c.y - 1}
	below := coord{c.x, c.y + 1}

	neighbours := make([]neighbour, 0, 4)
	if left.x >= 0 {
		neighbours = append(neighbours, neighbour{left, h.Value(left)})
	}

	if right.x < len(h.data[c.y]) {
		neighbours = append(neighbours, neighbour{right, h.Value(right)})
	}

	if above.y >= 0 {
		neighbours = append(neighbours, neighbour{above, h.Value(above)})
	}

	if below.y < len(h.data) {
		neighbours = append(neighbours, neighbour{below, h.Value(below)})
	}

	return neighbours
}

// Returns the side of the side of the heatmap, assumes the map is square.
func (h *HM) Size() (int, int) {
	rows := len(h.data)
	cols := len(h.data[0])
	return rows, cols
}

func (h *HM) Value(c coord) int {
	return h.data[c.y][c.x]
}

func (h *HM) SetValue(c coord, v int) {
	h.data[c.y][c.x] = v
}

func TestNeighbours(t *testing.T) {
	hm := &HM{[][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}}

	want := [][][]int{
		{{2, 4}, {1, 3, 5}, {2, 6}},
		{{5, 1, 7}, {4, 6, 2, 8}, {5, 3, 9}},
		{{8, 4}, {7, 9, 5}, {8, 6}},
	}

	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			for n, g := range hm.Neighbours(coord{x, y}) {
				t.Log(x, y, n, g)
				if g.v != want[y][x][n] {
					t.Fatal(g, want[y][x][n])
				}

			}
		}
	}
}

func ReadHeightMap(t *testing.T) *HM {
	rows, err := test.Read(test.String)
	if err != nil {
		t.Fatal(err)
	}

	data, err := parseHeightMap(rows)
	if err != nil {
		t.Fatal(err)
	}

	return &HM{data}
}

func parseHeightMap(rows []string) ([][]int, error) {
	data := [][]int{}
	for _, rt := range rows {
		ri := make([]int, len(rt))
		for j, dt := range rt {
			di, err := strconv.ParseInt(string(dt), 10, 32)
			if err != nil {
				return nil, err
			}

			ri[j] = int(di)
		}

		data = append(data, ri)
	}

	return data, nil
}

func TestParseHeightMap(t *testing.T) {
	rows := []string{
		"12345",
		"23456",
		"34567",
		"45678",
		"56789",
	}

	want := [][]int{
		{1, 2, 3, 4, 5},
		{2, 3, 4, 5, 6},
		{3, 4, 5, 6, 7},
		{4, 5, 6, 7, 8},
		{5, 6, 7, 8, 9},
	}

	got, _ := parseHeightMap(rows)

	if !reflect.DeepEqual(got, want) {
		t.Fatal(got, want)
	}
}

type coord struct {
	x int
	y int
}

func findLowpoints(hm *HM) []coord {
	lowpoints := []coord{}

	rows, cols := hm.Size()

	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			v := hm.Value(coord{x, y})
			isMin := false
			for _, n := range hm.Neighbours(coord{x, y}) {
				isMin = v < n.v
				if !isMin {
					break
				}
			}

			if isMin {
				lowpoints = append(lowpoints, coord{x, y})
			}
		}
	}

	return lowpoints
}

func TestDec9(t *testing.T) {
	test.Run(func() {
		hm := ReadHeightMap(t)
		lowpoints := findLowpoints(hm)

		sum := 0
		for _, lp := range lowpoints {
			sum += hm.Value(lp) + 1
		}

		test.Check(t, sum, 15, 486)
	})
}

func TestMapBasin(t *testing.T) {
	tests := []struct {
		txt   []string
		coord coord
		basin map[coord]int
	}{
		{
			[]string{
				"2199943210",
				"3987894921",
				"9856789892",
				"8767896789",
				"9899965678",
			},
			coord{6, 4},
			map[coord]int{
				{5, 4}: 6,
				{6, 3}: 6,
				{6, 4}: 5,
				{7, 2}: 8,
				{7, 3}: 7,
				{7, 4}: 6,
				{8, 3}: 8,
				{8, 4}: 7,
				{9, 4}: 8,
			},
		},
		{
			[]string{
				"99887678932398943999888667",
				"89776587891987899898765456",
				"76543456789996789789894347",
				"98432345899985345678932123",
				"87543467899874257899653999",
				"97665678901965346789869878",
				"98786989219875457896998767",
			},
			coord{23, 3},
			map[coord]int{
				{19, 1}: 8,
				{20, 0}: 8,
				{20, 1}: 7,
				{20, 2}: 8,
				{20, 4}: 6,
				{20, 5}: 8,
				{21, 0}: 8,
				{21, 1}: 6,
				{21, 3}: 3,
				{21, 4}: 5,
				{21, 5}: 6,
				{22, 0}: 8,
				{22, 1}: 5,
				{22, 2}: 4,
				{22, 3}: 2,
				{22, 4}: 3,
				{23, 0}: 6,
				{23, 1}: 4,
				{23, 2}: 3,
				{23, 3}: 1,
				{24, 0}: 6,
				{24, 1}: 5,
				{24, 2}: 4,
				{24, 3}: 2,
				{25, 0}: 7,
				{25, 1}: 6,
				{25, 2}: 7,
				{25, 3}: 3,
			},
		},
	}

	for _, tc := range tests {
		data, err := parseHeightMap(tc.txt)
		if err != nil {
			t.Fatal(err)
		}

		hm := &HM{data}

		basin := mapBasin(t, tc.coord, hm)
		logBasins(t, basin, hm)

		if !reflect.DeepEqual(tc.basin, basin) {
			t.Log("g:", basin)
			t.Fatal("w:", tc.basin)
		}
	}
}

func mapBasin(t *testing.T, c coord, hm *HM) map[coord]int {
	at := 0
	// The starting coordinate is part of the basin but it
	// is added as the starting point for the work to be done.
	// This allow us to keep the search generic and at the
	// cost of doint one extra work item.
	work := []neighbour{{c, hm.Value(c)}}
	basin := map[coord]int{}

	for at < len(work) {
		cur := work[at]
		at++
		// If the current location coordinate have been
		// added to the basin then we can move on directly.
		if v, ok := basin[cur.c]; ok {
			// Sanity check the value.
			if v != cur.v {
				panic("values must match")
			}
			continue
		}

		for _, n := range hm.Neighbours(cur.c) {
			if n.v < cur.v || n.v == 9 {
				continue
			}

			work = append(work, n)
		}

		basin[cur.c] = cur.v
	}

	return basin
}

func findBasinSize(t *testing.T, c coord, hm *HM) int {
	t.Log("basin", c)
	basin := mapBasin(t, c, hm)
	t.Log(">>>>>", basin)
	return len(basin)
}

func TestDec9Debug(t *testing.T) {
	test.Run(func() {
		hm := ReadHeightMap(t)
		lps := findLowpoints(hm)
		basins := map[coord]int{}

		for _, lp := range lps {
			basin := mapBasin(t, lp, hm)

			for c, v := range basin {
				if v != hm.Value(c) {
					t.Fatal("wrong value at:", c, " ", v, hm.Value(c))
				}
			}

			for c, v := range basin {
				if _, ok := basins[c]; ok {
					t.Fatal("overlapping basins", c)
				}

				basins[c] = v
			}
		}

		for r := 0; r < len(hm.data); r++ {
			rt := &strings.Builder{}
			rt.WriteString(fmt.Sprintf("%d: ", r))
			for c := 0; c < len(hm.data[r]); c++ {
				cur := coord{c, r}
				if _, ok := basins[cur]; ok {
					rt.WriteString("x")
				} else {
					rt.WriteString(fmt.Sprintf("%d", hm.Value(cur)))
				}
			}

			t.Log(rt.String())
		}
	})
}

func logBasins(t *testing.T, basins map[coord]int, hm *HM) {
	for r := 0; r < len(hm.data); r++ {
		rt := &strings.Builder{}
		rt.WriteString(fmt.Sprintf("%d: ", r))
		for c := 0; c < len(hm.data[r]); c++ {
			cur := coord{c, r}
			if _, ok := basins[cur]; ok {
				rt.WriteString("x")
			} else {
				rt.WriteString(fmt.Sprintf("%d", hm.Value(cur)))
			}
		}

		t.Log(rt.String())
	}
}

func TestDec9_2(t *testing.T) {
	test.Run(func() {
		hm := ReadHeightMap(t)
		lowpoints := findLowpoints(hm)
		sizes := []int{}
		for _, lp := range lowpoints {
			size := findBasinSize(t, lp, hm)
			sizes = append(sizes, size)
		}

		sort.Sort(sort.Reverse(sort.IntSlice(sizes)))
		product := 1
		for i := 0; i < 3; i++ {
			size := sizes[i]
			t.Log(size)
			product *= size
		}

		t.Log(sizes)
		test.Check(t, product, 1134, 1059300)
	})
}

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}
