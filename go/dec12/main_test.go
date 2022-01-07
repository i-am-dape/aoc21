package main

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/i-am-dape/aoc21/go/test"
)

type graph struct {
	nodes map[string]map[string]interface{}
}

func parseGraph(lines []string) graph {
	nodes := map[string]map[string]interface{}{}

	insert := func(from, to string) {
		nf, ok := nodes[from]
		if !ok {
			nf = map[string]interface{}{}
			nodes[from] = nf
		}

		nf[to] = 1
	}

	// This will add the basic connections but not all connections are
	// explicitly part of the input. If 'A-b' exists in the input then an edge
	// 'b-A' exists in the graph even if it isn't specified in the input.
	for _, l := range lines {
		ix := strings.IndexByte(l, '-')
		from := l[:ix]
		to := l[ix+1:]

		insert(from, to)
		insert(to, from)
	}

	return graph{nodes}
}

func TestParseGraph(t *testing.T) {
	w := graph{map[string]map[string]interface{}{
		"start": {
			"A": 1,
			"b": 1,
		},
		"A": {
			"start": 1,
			"c":     1,
			"b":     1,
			"end":   1,
		},
		"b": {
			"start": 1,
			"A":     1,
			"d":     1,
			"end":   1,
		},
		"c": {
			"A": 1,
		},
		"d": {
			"b": 1,
		},
		"end": {
			"A": 1,
			"b": 1,
		},
	}}
	g := parseGraph([]string{
		"start-A",
		"start-b",
		"A-c",
		"A-b",
		"b-d",
		"A-end",
		"b-end",
	})

	if !reflect.DeepEqual(g, w) {
		t.Log(g)
		t.Log(w)
		t.Fatal()
	}
}

func readGraph(t *testing.T) graph {
	l, err := test.Read(test.String)
	if err != nil {
		t.Fatal(err)
	}

	return parseGraph(l)
}

func TestDec12(t *testing.T) {
	test.Run(func() {
		g := readGraph(t)
		u := findUniquePathsInGraph(t, g)
		test.Check(t, u, 226, 3510)
	})
}

func peek(items []string) string {
	return items[len(items)-1]
}

func pop(items []string) []string {
	return items[0 : len(items)-1]
}

func push(items []string, item string) []string {
	return append(items, item)
}

func contains(item string, items []string) bool {
	for _, i := range items {
		if i == item {
			return true
		}
	}

	return false
}

func walk(n int, g graph, previous []string, completed map[string]interface{}) {
	fmt.Println(previous)
	parent := peek(previous)
	if parent == "end" {
		completed[fmt.Sprint(previous)] = 1
		return
	}

	for edge := range g.nodes[parent] {
		if strings.ToLower(edge) == edge && contains(edge, previous) {
			continue
		}

		previous = push(previous, edge)
		walk(n+1, g, previous, completed)
		previous = pop(previous)
	}
}

func TestDec12TinyExamplex(t *testing.T) {
	tests := []struct {
		lines  []string
		unique int
	}{
		{
			[]string{
				"start-A",
				"start-b",
				"A-c",
				"A-b",
				"b-d",
				"A-end",
				"b-end",
			},
			10,
		},
		{
			[]string{
				"dc-end",
				"HN-start",
				"start-kj",
				"dc-start",
				"dc-HN",
				"LN-dc",
				"HN-end",
				"kj-sa",
				"kj-HN",
				"kj-dc",
			},
			19,
		},
		{
			[]string{
				"fs-end",
				"he-DX",
				"fs-he",
				"start-DX",
				"pj-DX",
				"end-zg",
				"zg-sl",
				"zg-pj",
				"pj-he",
				"RW-he",
				"fs-DX",
				"pj-RW",
				"zg-RW",
				"start-pj",
				"he-WI",
				"zg-he",
				"pj-fs",
				"start-RW",
			},
			226,
		},
	}

	for _, tc := range tests {
		unique := findUniquePaths(t, tc.lines)
		if unique != tc.unique {
			t.Fatal(unique, tc.unique)
		}
	}
}

func findUniquePaths(t *testing.T, lines []string) int {
	g := parseGraph(lines)
	return findUniquePathsInGraph(t, g)
}

func findUniquePathsInGraph(t *testing.T, g graph) int {
	completed := map[string]interface{}{}
	walk(0, g, []string{"start"}, completed)

	t.Log(len(completed))
	t.Log(completed)
	return len(completed)
}

func TestDec12_2(t *testing.T) {
	test.Run(func() {
	})
}

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}
