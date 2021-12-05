package main

import (
	"flag"
	"os"
	"testing"

	"github.com/i-am-dape/aoc21/go/test"
)

func TestDec1(t *testing.T) {
	test.Run(func() {

		input, err := test.Read(test.Dec2Int)
		if err != nil {
			panic(err)
		}

		increase := 0
		prev := input[0]
		for i := 1; i < len(input); i++ {
			cur := input[i]
			if cur > prev {
				increase++
			}
			prev = cur
		}

		test.Check(t, increase, 7, 1676)
	})
}

func sum(in []int) int {
	s := 0
	for _, n := range in {
		s += n
	}
	return s
}

func TestDec1_2(t *testing.T) {
	test.Run(func() {
		input, err := test.Read(test.Dec2Int)
		if err != nil {
			panic(err)
		}

		incr := 0
		prev := sum([]int{input[0], input[1], input[2]})
		for i := 3; i < len(input); i++ {
			cur := sum([]int{input[i-2], input[i-1], input[i]})

			if cur > prev {
				incr++
			}

			prev = cur
		}

		test.Check(t, incr, 5, 1706)
	})
}

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}
