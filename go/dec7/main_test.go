package main

import (
	"flag"
	"fmt"
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

		t.Log("tot:", total)

		t.Log(cc[16][2])
		t.Log(cc[1][2])
		t.Log(cc[2][2])
		t.Log(cc[0][2])
		t.Log(cc[4][2])
		t.Log(cc[7][2])

		distances := map[int]int{}

		for i := min; i < max; i++ {
			fmt.Println(i)
			totalDistance := 0
			for pos, cnt := range bc {
				delta := int(pos) - i
				if delta < 0 {
					delta *= -1
				}

				total := delta * cnt
				fmt.Println(i, ")", pos, delta, cnt, total)

				totalDistance += total
			}

			distances[i] = totalDistance
		}

		fmt.Println(distances)

		minPos := -1
		minDist := math.MaxInt
		for pos, dist := range distances {
			if dist < minDist {
				minDist = dist
				minPos = pos
			}
		}

		fmt.Println(minPos, minDist)
	})
}

func TestDec7_2(t *testing.T) {
	test.Run(func() {
	})
}

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}
