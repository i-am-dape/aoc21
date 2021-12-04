package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/i-am-dape/aoc21/go/input"
)

type cmd string

func (c cmd) IsForward() bool {
	return c.hasPrefix("forward")
}

func (c cmd) hasPrefix(prefix string) bool {
	return strings.HasPrefix(string(c), prefix)
}

func (c cmd) parseIntAt(at int) int {
	n, err := strconv.ParseInt(string(c)[at:], 10, 32)
	if err != nil {
		panic(err)
	}

	return int(n)
}

func (c cmd) Forward() int {
	return c.parseIntAt(8)
}

func (c cmd) IsUp() bool {
	return c.hasPrefix("up")
}

func (c cmd) Up() int {
	return c.parseIntAt(3)
}

func (c cmd) IsDown() bool {
	return c.hasPrefix("down")
}

func (c cmd) Down() int {
	return c.parseIntAt(5)
}

func main() {
	flag.Parse()
	input, err := input.Read(func(txt string) (cmd, error) {
		return cmd(txt), nil
	})
	if err != nil {
		panic(err)
	}

	horizontal := 0
	depth := 0

	for _, cmd := range input {
		switch {
		case cmd.IsForward():
			horizontal += cmd.Forward()
		case cmd.IsUp():
			depth -= cmd.Up()
		case cmd.IsDown():
			depth += cmd.Down()
		}
	}

	fmt.Println(horizontal, depth, horizontal*depth)
}
