package main

import (
	"flag"
	"os"
	"testing"

	"github.com/i-am-dape/aoc21/go/test"
)

func Test1(t *testing.T) {
	input, err := test.Read(test.Dec2Int)
	if err != nil {
		panic(err)
	}
	t.Log(input)
}

func Test2(t *testing.T) {
}

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}
