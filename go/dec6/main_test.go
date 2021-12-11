package main

import (
	"flag"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/i-am-dape/aoc21/go/test"
)

var NoFish = Fish(-1)

type Fish int

func (f *Fish) Tick() (Fish, bool) {
	n := NoFish
	*f -= 1

	if *f < 0 {
		*f = 6
		n = Fish(8)
	}

	return n, n != NoFish
}

func TestTick(t *testing.T) {
	f := Fish(3)

	// Codified version of the description from adventofcode.com/2021/day/6

	tests := []struct {
		next Fish
		n    Fish
		ok   bool
	}{
		{Fish(2), NoFish, false},
		{Fish(1), NoFish, false},
		{Fish(0), NoFish, false},
		{Fish(6), Fish(8), true},
	}

	for _, tc := range tests {
		n, ok := f.Tick()
		t.Log(f, tc)

		if f != tc.next || n != tc.n || ok != tc.ok {
			t.Fatal()
		}
	}
}

func Read(t *testing.T) []Fish {
	lines, err := test.Read(test.String)
	if err != nil {
		t.Fatal(err)
	}

	txts := strings.Split(lines[0], ",")
	fishes := []Fish{}
	for _, txt := range txts {
		n, err := strconv.ParseInt(txt, 10, 32)
		if err != nil {
			t.Fatal(err)
		}

		fishes = append(fishes, Fish(n))
	}

	return fishes
}

func simulate(t *testing.T, days int, fishes []Fish) []Fish {
	for i := 0; i < days; i++ {
		t.Logf("Day:%d %d", i, len(fishes))

		for at, f := range fishes {
			n, ok := f.Tick()
			fishes[at] = f
			if ok {
				fishes = append(fishes, n)
			}
		}
	}

	return fishes
}

func simulateFor(t *testing.T, days int, example, input int64) {
	test.Run(func() {
		fs := Read(t)
		t.Log(fs)
		res := simulate(t, days, fs)

		test.Check64(t, int64(len(res)), example, input)
	})
}

func TestDec6(t *testing.T) {
	simulateFor(t, 80, 5934, 351092)
}

func TestDec6_2(t *testing.T) {
	simulateFor(t, 256, 26984457539, 0)
}

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}
