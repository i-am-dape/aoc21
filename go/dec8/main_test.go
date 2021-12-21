package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/i-am-dape/aoc21/go/test"
)

type input struct {
	signals []string
	output  []string
}

func string2Input(txt string) (input, error) {
	parts := strings.Split(txt, "|")
	split := func(txt string) []string {
		tmp := strings.Split(txt, " ")
		wat := 0
		for i, _ := range tmp {
			if len(tmp[i]) == 0 {
				continue
			}
			//fmt.Println(tmp[i], len(tmp[i]))
			tmp[wat] = strings.TrimSpace(tmp[i])
			wat++
		}
		return tmp[:wat]
	}

	return input{split(parts[0]), split(parts[1])}, nil
}

var exampleInputTxt = "acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf"
var exampleInput = input{
	[]string{"acedgfb", "cdfbe", "gcdfa", "fbcad", "dab", "cefabd", "cdfgeb", "eafb", "cagedb", "ab"},
	[]string{"cdfeb", "fcadb", "cdfeb", "cdbaf"},
}

func TestString2Input(t *testing.T) {
	g, _ := string2Input(exampleInputTxt)
	w := exampleInput

	if !reflect.DeepEqual(g, w) {
		t.Log(g)
		t.Log(w)
		t.Fatal()
	}
}

var segments = []rune{
	'a', 'b', 'c', 'd', 'e', 'f', 'g',
}

var digits = []string{
	"abcefg",
	"cf",
	"acdeg",
	"acdfg",
	"bcdf",
	"abdfg",
	"abdefg",
	"acf",
	"abcdefg",
	"abcdfg",
}

var segment2count = map[string]int{}

func init() {
	segment2count[digits[0]] = len(digits[0])
	segment2count[digits[1]] = len(digits[1])
	segment2count[digits[2]] = len(digits[2])
	segment2count[digits[3]] = len(digits[3])
	segment2count[digits[4]] = len(digits[4])
	segment2count[digits[5]] = len(digits[5])
	segment2count[digits[6]] = len(digits[6])
	segment2count[digits[7]] = len(digits[7])
	segment2count[digits[8]] = len(digits[8])
	segment2count[digits[9]] = len(digits[9])
}

func TestDec8(t *testing.T) {
	test.Run(func() {
		in, err := test.Read(string2Input)
		if err != nil {
			t.Fatal(err)
		}

		count := 0
		for _, input := range in {
			for _, txt := range input.output {
				ln := len(txt)
				if ln == 2 || ln == 4 || ln == 3 || ln == 7 {
					count++
				}
			}
		}

		test.Check(t, count, 26, 294)
	})
}

func bucketBySize(signals []string) map[int][]string {
	buckets := map[int][]string{}
	for _, signal := range signals {
		size := len(signal)
		bucket, ok := buckets[size]
		if !ok {
			bucket = []string{}
		}

		bucket = append(bucket, signal)
		buckets[size] = bucket
	}
	return buckets
}

func findMissing(short, long string) string {
	if len(long) != len(short)+1 {
		panic("bad size")
	}

	for _, c := range long {
		if strings.IndexRune(short, c) < 0 {
			return string(c)
		}
	}

	panic("unreachable")
}

func TestFindMissing(t *testing.T) {
	tests := []struct {
		txt1    string
		txt2    string
		missing string
	}{
		{"acf", "cf", "a"},
		{"bdg", "gd", "b"},
	}

	for _, tc := range tests {
		t.Log(tc)
		missing := findMissing(tc.txt2, tc.txt1)
		if missing != tc.missing {
			t.Fatal(tc.txt2, tc.txt1, missing)
		}

		stxt2 := shuffleString(tc.txt2)
		stxt1 := shuffleString(tc.txt1)
		smissing := findMissing(stxt2, stxt1)
		if smissing != tc.missing {
			t.Fatal(stxt2, stxt1, smissing)
		}
	}
}

func findSegmentA(bySize map[int][]string) string {
	d7 := bySize[3][0]
	d1 := bySize[2][0]

	return findMissing(d1, d7)
}

func findSegmentD(bySize map[int][]string) string {
	d235 := bySize[5]
	d4 := bySize[4]

	buf := &strings.Builder{}
	for _, d := range d235 {
		buf.WriteString(d)
	}

	buf.WriteString(d4[0])

	cnt := countUnique(buf.String())
	for r, c := range cnt {
		if c == 4 {
			return string(r)
		}
	}

	panic("unreachable")
}

func findSegmentE(bySize map[int][]string, sA, sG string) string {
	d4 := bySize[4][0]
	d8 := bySize[7][0]

	buf := &strings.Builder{}
	buf.WriteString(d4)
	buf.WriteString(d8)
	buf.WriteString(sA)
	buf.WriteString(sG)

	cnt := countUnique(buf.String())
	for r, c := range cnt {
		if c == 1 {
			return string(r)
		}
	}

	panic("unreachable")
}

func countUnique(txt string) map[rune]int {
	cnts := map[rune]int{}
	for _, r := range txt {
		cnts[r]++
	}
	return cnts
}

func findSegmentG(bySize map[int][]string, sA string) string {
	d235 := bySize[5]
	d069 := bySize[6]

	buf := &strings.Builder{}
	for _, d := range append(d235, d069...) {
		buf.WriteString(d)
	}

	cnt := countUnique(strings.ReplaceAll(buf.String(), sA, ""))
	for r, c := range cnt {
		if c == 6 {
			return string(r)
		}
	}

	panic("unreachable")
}

func findSegmentB(bySize map[int][]string, sD, sE, sG string) string {
	d7 := bySize[3]
	d8 := bySize[7]

	buf := &strings.Builder{}
	for _, d := range append(d7, d8...) {
		buf.WriteString(d)
	}

	buf.WriteString(sD)
	buf.WriteString(sE)
	buf.WriteString(sG)

	cnt := countUnique(buf.String())
	fmt.Println(cnt, segments)
	for c, n := range cnt {
		if n == 1 {
			return string(c)
		}
	}

	panic("unreachable")
}

func runeMap(txt string) map[rune]int {
	m := map[rune]int{}
	for _, r := range txt {
		m[r] = 0
	}

	return m
}

func findSegmentC(bySize map[int][]string, sD string, sE string) string {
	d069 := bySize[6]

	buf := &strings.Builder{}
	for _, d := range d069 {
		buf.WriteString(d)
	}

	buf.WriteString(sD)
	buf.WriteString(sE)

	cnt := countUnique(buf.String())
	fmt.Println(cnt)
	for r, c := range cnt {
		if c == 2 {
			return string(r)
		}
	}

	panic("unreachable")
}

func findSegmentF(sA, sB, sC, sD, sE, sG string) string {
	buf := &strings.Builder{}
	buf.WriteString(sA)
	buf.WriteString(sB)
	buf.WriteString(sC)
	buf.WriteString(sD)
	buf.WriteString(sE)
	buf.WriteString(sG)

	used := buf.String()
	all := "abcdefg"

	for _, r := range all {
		if strings.IndexRune(used, r) < 0 {
			return string(r)
		}
	}

	panic("unreachable")
}

func deduct(signals []string) map[string]string {
	bySize := bucketBySize(signals)

	for i := 0; i < 10; i++ {
		fmt.Println(i, bySize[i])
	}

	sA := findSegmentA(bySize)
	sD := findSegmentD(bySize)
	sG := findSegmentG(bySize, sA)
	sE := findSegmentE(bySize, sA, sG)
	sB := findSegmentB(bySize, sD, sE, sG)
	sC := findSegmentC(bySize, sD, sE)
	sF := findSegmentF(sA, sB, sC, sD, sE, sG)

	fmt.Println("a", sA)
	fmt.Println("b", sB)
	fmt.Println("c", sC)
	fmt.Println("d", sD)
	fmt.Println("e", sE)
	fmt.Println("f", sF)
	fmt.Println("g", sG)

	//return map[string]string{"a": sA, "b": sB, "c": sC, "d": sD, "e": sE, "f": sF, "g": sG}
	return map[string]string{sA: "a", sB: "b", sC: "c", sD: "d", sE: "e", sF: "f", sG: "g"}
}

func shuffleString(txt string) string {
	buf := make([]byte, len(txt))

	randIntn := func(used map[int]interface{}) int {
		for {
			dest := rand.Intn(len(txt))
			if _, ok := used[dest]; !ok {
				used[dest] = struct{}{}
				return dest
			}
		}
	}

	usedDest := map[int]interface{}{}
	usedSrc := map[int]interface{}{}

	for i := 0; i < len(txt); i++ {
		buf[randIntn(usedDest)] = txt[randIntn(usedSrc)]
	}

	return string(buf)
}

func TestDisplay(t *testing.T) {
	for n, d := range digits {
		got := signals2int(d)
		t.Log(d, "got:", got, "want:", n)
		if got != n {
			t.Fatal()
		}

		sd := shuffleString(d)
		sgot := signals2int(sd)
		t.Log(sd, "got:", sgot, "want:", n)
		if got != n {
			t.Fatal()
		}
	}
}

func signals2int(signals string) int {
	for n, digit := range digits {
		if len(digit) != len(signals) {
			continue
		}

		count := 0
		for _, c := range signals {
			if strings.IndexRune(digit, c) < 0 {
				break
			}

			count++
		}

		if count == len(digit) {
			return n
		}
	}

	panic("unable to resolve signals")
}

func display(conn map[string]string, signals string) int {
	sb := &strings.Builder{}
	for _, r := range signals {
		rev := conn[string(r)]
		sb.WriteString(rev)
	}

	translated := sb.String()
	return signals2int(translated)
}

func TestDeduct(t *testing.T) {
	fmt.Printf("%s : %s", exampleInput.signals, exampleInput.output)
	fmt.Println()
	fmt.Println(digits)
	conn := deduct(exampleInput.signals)

	revConn := map[string]string{}
	for k, v := range conn {
		revConn[v] = k
	}

	checkRevConn := func(signal, w string) {
		g := revConn[signal]
		if g != w {
			t.Fatal(g, w)
		}
	}

	checkRevConn("a", "d")
	checkRevConn("b", "e")
	checkRevConn("c", "a")
	checkRevConn("d", "f")
	checkRevConn("e", "g")
	checkRevConn("f", "b")
	checkRevConn("g", "c")
}

func TestDeductRandom(t *testing.T) {
	// Returns ten digits from a shuffled 7-segment display.
	shuffleSegments := func() []string {
		signals := "abcdefg"
		shuffled := shuffleString(signals)
		t.Log(signals, "->", shuffled)
		lookup := map[rune]rune{}
		for i, signal := range signals {
			lookup[signal] = rune(shuffled[i])
		}
		t.Log(lookup)

		res := []string{}
		for i := 0; i < 10; i++ {
			signal := digits[i]
			sb := &strings.Builder{}
			for _, s := range signal {
				sb.WriteRune(lookup[s])
			}
			res = append(res, sb.String())
			t.Log(signal, "->", sb.String())
		}

		return res
	}

	for i := 0; i < 1; i++ {
		signals := shuffleSegments()
		conn := deduct(signals)
		t.Log(conn)

		signals2 := shuffleSegments()
		conn2 := deduct(signals2)
		t.Log(conn2)
	}
}

func processInput(i input) int {
	conn := deduct(i.signals)
	fmt.Println("connections map", conn)

	got := []int{}
	for _, code := range i.output {
		n := display(conn, code)
		fmt.Println(code, n)
		got = append(got, n)
	}

	//for _, code := range i.signals {
	//	fmt.Println(code, display(conn, code))
	//}

	total := (1000 * got[0]) + (100 * got[1]) + (10 * got[2]) + got[3]
	return total
}

func TestDec8_2FixedExample(t *testing.T) {
	fmt.Printf("%s : %s", exampleInput.signals, exampleInput.output)
	fmt.Println()
	fmt.Println(digits)

	total := processInput(exampleInput)
	want := 5353
	if total != want {
		t.Fatal(total, want)
	}
}

func TestDec8_2(t *testing.T) {
	test.Run(func() {
		inputs, err := test.Read(string2Input)
		if err != nil {
			t.Fatal(err)
		}

		total := 0
		for _, input := range inputs {
			t.Log(input)
			step := processInput(input)
			t.Log(step)
			total += step
		}

		test.Check(t, total, 61229, 973292)
	})
}

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}
