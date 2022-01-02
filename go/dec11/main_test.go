package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/i-am-dape/aoc21/go/test"
)

type grid struct {
	data [][]int
}

func (g *grid) String() string {
	sb := &strings.Builder{}
	for r := range g.data {
		for c := range g.data[r] {
			sb.WriteByte(byte(48 + g.data[r][c]))
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

type coord struct {
	r, c int
}

func (g *grid) step() int {
	flashed := map[coord]interface{}{}
	for r := range g.data {
		for c := range g.data[r] {
			g.data[r][c] += 1
		}
	}

	count := math.MaxInt
	i := 0
	for count > 0 {
		i++
		count = 0
		for r := range g.data {
			for c := range g.data[r] {
				cur := coord{r, c}
				if _, ok := flashed[cur]; ok {
					continue
				}

				if g.data[r][c] > 9 {
					//fmt.Print("flash [", r, c, "] >> ")
					count++
					flashed[coord{r, c}] = 1

					for nr := r - 1; nr <= r+1; nr++ {
						for nc := c - 1; nc <= c+1; nc++ {
							if nr < 0 || nr >= len(g.data) || nc < 0 || nc >= len(g.data[0]) {
								continue
							}

							if nr == r && nc == c {
								continue
							}

							//fmt.Print("[", nr, nc, "]")
							g.data[nr][nc]++
						}
					}

					//fmt.Println()
				}
			}
		}
	}

	for c, _ := range flashed {
		g.data[c.r][c.c] = 0
	}

	return len(flashed)
}

func parseLine(l string) []int {
	vals := []int{}
	for _, e := range l {
		vals = append(vals, int(e)-48)
	}
	return vals
}

func TestParseLine(t *testing.T) {
	got := parseLine("5264556173")
	want := []int{5, 2, 6, 4, 5, 5, 6, 1, 7, 3}

	if !reflect.DeepEqual(got, want) {
		t.Log(got)
		t.Log(want)
		t.Fatal()
	}
}

func readGrid(t *testing.T) grid {
	lines, err := test.Read(test.String)
	if err != nil {
		t.Fatal(err)
	}

	g := grid{}
	for _, l := range lines {
		vals := []int{}
		for _, e := range l {
			vals = append(vals, int(e)-48)
		}
		g.data = append(g.data, vals)
	}

	return g
}

func TestDec11(t *testing.T) {
	test.Run(func() {
		grid := readGrid(t)
		fmt.Printf("%s", &grid)

		flashes := 0
		for i := 0; i < 100; i++ {
			flashes += grid.step()
			fmt.Println("\n", i, flashes)
			fmt.Printf("%s", &grid)
		}

		test.Check(t, flashes, 1656, 1591)
	})
}

func TestDec11_2(t *testing.T) {
	test.Run(func() {
		grid := readGrid(t)
		fmt.Printf("%s", &grid)

		flashed := 0
		step := 0
		for flashed != 100 {
			flashed = grid.step()
			step++
			fmt.Println("\n", step, flashed)
			fmt.Printf("%s", &grid)
		}

		test.Check(t, step, 195, 1591)
	})
}

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}
