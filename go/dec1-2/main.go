package main

import (
	"flag"
	"fmt"

	"github.com/i-am-dape/aoc21/go/input"
)

func sum(in []int) int {
	s := 0
	for _, n := range in {
		s += n
	}
	return s
}

func main() {
	flag.Parse()
	input, err := input.Read(input.Dec2Int)
	if err != nil {
		panic(err)
	}

	incr := 0
	prev := sum([]int{input[0], input[1], input[2]})
	fmt.Println(prev)
	for i := 3; i < len(input); i++ {
		cur := sum([]int{input[i-2], input[i-1], input[i]})
		fmt.Println(cur)

		if cur > prev {
			incr++
		}

		prev = cur
	}
	fmt.Println(incr)
}
