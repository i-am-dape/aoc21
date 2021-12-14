package main

import (
	"flag"
	"fmt"
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

func TestDec1KB(t *testing.T) {
	// example := []int{
	// 	199,
	// 	200,
	// 	208,
	// 	210,
	// 	200,
	// 	207,
	// 	240,
	// 	269,
	// 	260,
	// 	263,
	// }
	test.Run(func() {

		example, err := test.Read(test.Dec2Int)
		if err != nil {
			panic(err)
		}

		count := 0
		for i := 1; i < len(example); i++ {
			if example[i] > example[i-1] {
				count++
			}
		}

		fmt.Print(count)

	})
}
