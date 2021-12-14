package main

import (
	"flag"
	"math"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/i-am-dape/aoc21/go/test"
)

func ReadCrabs(t *testing.T) []Crab {
	lines, err := test.Read(test.String)
	if err != nil {
		t.Fatal(err)
	}

	txts := strings.Split(lines[0], ",")
	crabs := []Crab{}
	for _, txt := range txts {
		n, err := strconv.ParseInt(txt, 10, 32)
		if err != nil {
			t.Fatal(err)
		}

		crabs = append(crabs, Crab(n))
	}

	return crabs
}

type Crab int

// Bucket adds up all crabs at the same level. (returns min, max, buckets)
func bucket(crabs []Crab) (int, int, map[Crab]int) {
	min := math.MaxInt
	max := math.MinInt

	cnt := map[Crab]int{}
	for _, crab := range crabs {
		if cur, ok := cnt[crab]; !ok {
			cnt[crab] = 1
		} else {
			cnt[crab] = cur + 1
		}

		if int(crab) < min {
			min = int(crab)
		}

		if int(crab) > max {
			max = int(crab)
		}
	}
	return min, max, cnt
}

func cross(bc map[Crab]int) map[Crab]map[Crab]int {
	res := map[Crab]map[Crab]int{}
	for src, cnt := range bc {
		dests := map[Crab]int{}
		res[src] = dests

		for dest, _ := range bc {
			if src == dest {
				continue
			}

			delta := int(src) - int(dest)
			if delta < 0 {
				delta = delta * -1
			}

			dests[dest] = delta * cnt
		}
	}

	return res
}

func TestDec7HardcodedExample(t *testing.T) {
	test.Run(func() {
		crabs := ReadCrabs(t)
		t.Log(crabs)
		min, max, bc := bucket(crabs)
		t.Log(min, max, len(bc))
		//t.Log(bc)
		cc := cross(bc)
		t.Log(len(cc))
		//t.Log(cc)

		ex := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
		exu := map[int]interface{}{}
		for _, x := range ex {
			exu[x] = nil
		}

		total := 0
		for src, _ := range exu {
			total += cc[Crab(src)][2]
		}
	})
}

func TestDec7(t *testing.T) {
	optimize(t, 0, 0, func(dist int) int { return dist })
}

func dist2fuel(dist int) int {
	total := 0

	for i := 1; i <= dist; i++ {
		total += i
	}

	return total
}

func TestDec7_dist2fuel(t *testing.T) {
	tests := []struct {
		n    int
		cost int
	}{
		{1, 1},
		{2, 3},
		{3, 6},
		{11, 66},
		{4, 10},
	}

	for _, tc := range tests {
		t.Log(tc)

		cost := dist2fuel(tc.n)
		if cost != tc.cost {
			t.Error(">", tc.n, cost, tc.cost)
		}
	}
}

func TestDec7_2(t *testing.T) {
	optimize(t, 168, 93397632, dist2fuel)
}

func optimize(t *testing.T, example, input int, fuelFunc func(int) int) {
	test.Run(func() {
		crabs := ReadCrabs(t)
		//t.Log(crabs)
		min, max, bc := bucket(crabs)
		//t.Log(min, max, len(bc))
		distances := map[int]int{}

		for i := min; i < max; i++ {
			totalDistance := 0
			for pos, cnt := range bc {
				delta := int(pos) - i
				if delta < 0 {
					delta *= -1
				}

				total := fuelFunc(delta) * cnt

				totalDistance += total
			}

			distances[i] = totalDistance
		}

		minDist := math.MaxInt
		for _, dist := range distances {
			if dist < minDist {
				minDist = dist
			}
		}

		test.Check(t, minDist, example, input)
	})
}

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}
