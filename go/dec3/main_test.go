package main

import (
	"flag"
	"os"
	"strconv"
	"testing"

	"github.com/i-am-dape/aoc21/go/test"
)

func countValues(reports []string, at int) (o int, z int) {
	for _, report := range reports {
		if report[at] == '0' {
			z++
		} else {
			o++
		}
	}
	return o, z
}

func findMostCommonValue(reports []string, at int, ifEq int) int {
	one, zero := countValues(reports, at)

	if one == zero {
		return ifEq
	}

	res := 0
	if one > zero {
		res = 1
	}
	return res
}

func gamma(reports []string, length int) int {
	gamma := 0

	for i := 0; i < length; i++ {
		gamma = gamma << 1
		mcv := findMostCommonValue(reports, i, -1)
		gamma = gamma | mcv
	}

	return gamma
}

func TestDec3(t *testing.T) {
	test.Run(func() {
		input, err := test.Read(test.String)
		if err != nil {
			t.Fatal(err)
		}

		// all reports have the same length but we don't
		// know what it is; so we peek at the first row
		// to figure it out.
		l := len(input[0])

		g := gamma(input, l)
		mask := 0
		for i := 0; i < l; i++ {
			mask = mask << 1
			mask = mask | 1
		}

		e := g ^ mask
		test.Check(t, g*e, 198, 1307354)
	})
}

func filterReports(reports []string, at int, keep byte) []string {
	filtered := make([]string, 0, len(reports))
	for _, report := range reports {
		if report[at] == keep {
			filtered = append(filtered, report)
		}
	}
	return filtered
}

func findRating(reports []string, l int, bitCriteria func(int) byte) int {
	rem := reports

	for i := 0; i < l; i++ {
		mcv := findMostCommonValue(rem, i, 1)
		keep := bitCriteria(mcv)

		rem = filterReports(rem, i, keep)

		if len(rem) == 1 {
			n, err := strconv.ParseInt(rem[0], 2, 32)
			if err != nil {
				panic(err)
			}
			return int(n)
		}
	}

	panic("unexpected")
}

func oxGenRating(reports []string, l int) int {
	return findRating(reports, l, func(mcv int) byte {
		keep := byte('0')
		if mcv == 1 {
			keep = byte('1')
		}
		return keep
	})
}

func co2ScrubberRating(reports []string, l int) int {
	return findRating(reports, l, func(mcv int) byte {
		keep := byte('1')
		if mcv == 1 {
			keep = byte('0')
		}

		return keep
	})
}

func TestDec3_2(t *testing.T) {
	test.Run(func() {
		input, err := test.Read(test.String)
		if err != nil {
			t.Fatal(err)
		}

		// all reports have the same length but we don't
		// know what it is; so we peek at the first row
		// to figure it out.
		l := len(input[0])

		ox := oxGenRating(input, l)
		co2 := co2ScrubberRating(input, l)
		test.Check(t, ox*co2, 230, 482500)
	})
}

// TestParseInt verifies that the result is padded with zeros as expected.
func TestParseInt(t *testing.T) {
	tests := []struct {
		txt string
		val int64
	}{
		{"11111", 0b11111},
		{"00011111", 0b11111},
	}

	for _, tc := range tests {
		n, err := strconv.ParseInt(tc.txt, 2, 32)
		if err != nil {
			t.Fatal(err)
		}

		if n != tc.val {
			t.Fatal(tc.txt, n, tc.val)
		}
	}
}

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}
