package main

import (
	"flag"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/i-am-dape/aoc21/go/test"
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

func TestDec2(t *testing.T) {
	test.Run(func() {
		input, err := test.Read(func(txt string) (cmd, error) {
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

		test.Check(t, horizontal*depth, 150, 2039256)
	})
}

func TestDec2_2(t *testing.T) {
	test.Run(func() {
		input, err := test.Read(func(txt string) (cmd, error) {
			return cmd(txt), nil
		})
		if err != nil {
			panic(err)
		}

		horizontal := 0
		depth := 0
		aim := 0

		for _, cmd := range input {
			switch {
			case cmd.IsForward():
				fwd := cmd.Forward()
				horizontal += fwd
				deltaDepth := aim * fwd
				depth += deltaDepth

			case cmd.IsUp():
				aim -= cmd.Up()

			case cmd.IsDown():
				aim += cmd.Down()
			}
		}

		test.Check(t, horizontal*depth, 900, 1856459736)
	})
}

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}
