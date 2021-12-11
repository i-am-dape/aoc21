package test

import (
	"flag"
)

var (
	input = flag.String("input", "", "set to run a single test")
)

func Run(tf func()) {
	prev := filename
	defer func() {
		filename = prev
	}()

	if len(*input) == 0 {
		tf()
		filename = "input.txt"
		tf()
	} else {
		filename = *input
		tf()
	}
}
