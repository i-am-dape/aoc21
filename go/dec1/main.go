package main

import (
	"flag"
	"fmt"

	"github.com/i-am-dape/aoc21/go/input"
)

func main() {
	flag.Parse()
	input, err := input.Read(input.Dec2Int)
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

	fmt.Println(increase)
}
