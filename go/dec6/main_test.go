package main

import (
	"flag"
	"fmt"
	"os"
	"reflect"
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

type FishShoal struct {
	Fishes [9]int64
}

func (f *FishShoal) String() string {
	b := &strings.Builder{}
	b.WriteString(fmt.Sprintf(" 8: %d\n", f.Fishes[8]))
	b.WriteString(fmt.Sprintf(" 7: %d\n", f.Fishes[7]))
	b.WriteString(fmt.Sprintf(" 6: %d\n", f.Fishes[6]))
	b.WriteString(fmt.Sprintf(" 5: %d\n", f.Fishes[5]))
	b.WriteString(fmt.Sprintf(" 4: %d\n", f.Fishes[4]))
	b.WriteString(fmt.Sprintf(" 3: %d\n", f.Fishes[3]))
	b.WriteString(fmt.Sprintf(" 2: %d\n", f.Fishes[2]))
	b.WriteString(fmt.Sprintf(" 1: %d\n", f.Fishes[1]))
	b.WriteString(fmt.Sprintf(" 0: %d\n", f.Fishes[0]))
	return b.String()
}

func (f *FishShoal) Tick() {
	next := [9]int64{}
	for n, cur := range f.Fishes {
		dest := n - 1
		if dest < 0 {
			dest = 6
			next[8] = cur
		}

		next[dest] += cur
	}
	f.Fishes = next
}

func (f *FishShoal) Count() int64 {
	cnt := int64(0)
	for _, cur := range f.Fishes {
		cnt += cur
	}
	return cnt
}

func TestFishShoalTick(t *testing.T) {
	shoal := FishShoal{[9]int64{0, 0, 0, 1, 0, 0, 0, 0, 0}}
	t.Log("\n", shoal.String())

	// Codified version of the description from adventofcode.com/2021/day/6

	tests := []struct {
		state [9]int64
	}{
		{[9]int64{0, 0, 1, 0, 0, 0, 0, 0, 0}},
		{[9]int64{0, 1, 0, 0, 0, 0, 0, 0, 0}},
		{[9]int64{1, 0, 0, 0, 0, 0, 0, 0, 0}},
		{[9]int64{0, 0, 0, 0, 0, 0, 1, 0, 1}},
		{[9]int64{0, 0, 0, 0, 0, 1, 0, 1, 0}},
	}

	for n, tc := range tests {
		shoal.Tick()
		t.Log("g:", n+1, shoal.Fishes)
		t.Log("w:", n+1, tc.state)

		//t.Log("\n", shoal.String())

		if !reflect.DeepEqual(shoal.Fishes, tc.state) {
			t.Fatal()
		}
	}
}

func simulateWithFishShoal(t *testing.T, days int, fishes []Fish, example, input int64) {
	shoal := FishShoal{[9]int64{}}
	for _, cur := range fishes {
		shoal.Fishes[cur]++
	}

	//fmt.Println(shoal.String())
	if int64(len(fishes)) != shoal.Count() {
		t.Fatal(len(fishes), shoal.Count())
	}

	t.Logf("Read: %d fishes", shoal.Count())

	for i := 0; i < days; i++ {
		shoal.Tick()

		if i == 17 {
			t.Log(shoal.Count())
		}
	}

	test.Check64(t, shoal.Count(), example, input)
}

func TestDec6WithFishShoal(t *testing.T) {
	test.Run(func() {
		simulateWithFishShoal(t, 80, Read(t), 5934, 351092)
	})
}

func TestDec6_2(t *testing.T) {
	test.Run(func() {
		simulateWithFishShoal(t, 256, Read(t), 26984457539, 1595330616005)
	})
}

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}
