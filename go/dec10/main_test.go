package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
	"testing"

	"github.com/i-am-dape/aoc21/go/test"
)

type chunk struct {
	open, close byte
	chunks      []chunk
	completed   bool
}

func (c chunk) String() string {
	sb := &strings.Builder{}
	sb.WriteByte(c.open)
	for _, c := range c.chunks {
		sb.WriteString(c.String())
	}
	sb.WriteByte(c.close)
	return sb.String()
}

var delims = map[byte]byte{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

type illegalCharError struct {
	line line
	msg  string
	got  byte
	want string
}

func (i illegalCharError) Error() string {
	return fmt.Sprintf("%s: %q got:%s want:%s", i.msg, i.line, string(i.got), i.want)
}

type unexpectedEndError struct {
	line line
}

func (u unexpectedEndError) Error() string {
	return fmt.Sprint("unexpected end", u.line)
}

// Parses a single chunk from the line. Implemented as a recursive
// decending parser, state is held in the shared line.
func parseChunk(l *line, complete bool) (chunk, error) {
	cur, ok := l.peek()
	if !ok {
		return chunk{}, unexpectedEndError{*l}
	}

	close, ok := delims[cur]
	if !ok {
		sb := &strings.Builder{}
		sep := ""
		for k := range delims {
			sb.WriteByte(k)
			sb.WriteString(sep)
			sep = ", "
		}
		return chunk{}, illegalCharError{*l, "expected open delimiter", cur, sb.String()}
	}

	l.pop()
	c := chunk{cur, close, nil, false}

	for {
		next, ok := l.peek()
		if !ok && complete {
			c.completed = true
			return c, nil
		} else if !ok {
			return chunk{}, unexpectedEndError{*l}
		}

		if next == close {
			l.pop()
			return c, nil
		}

		child, err := parseChunk(l, complete)
		if err != nil {
			return chunk{}, err
		}

		c.chunks = append(c.chunks, child)
	}
}

func TestParseLineComplete(t *testing.T) {
	tests := []struct {
		line string
		want string
	}{
		{
			"[({(<(())[]>[[{[]{<()<>>",
			"[({(<(())[]>[[{[]{<()<>>}}]])})]",
		},
		{
			"[(()[<>])]({[<{<<[]>>(",
			"[(()[<>])]({[<{<<[]>>()}>]})",
		},
		{
			"(((({<>}<{<{<>}{[]{[]{}",
			"(((({<>}<{<{<>}{[]{[]{}}}>}>))))",
		},
		{
			"{<[[]]>}<{[{[{[]{()[[[]",
			"{<[[]]>}<{[{[{[]{()[[[]]]}}]}]}>",
		},
		{
			"<{([{{}}[<[[[<>{}]]]>[]]",
			"<{([{{}}[<[[[<>{}]]]>[]]])}>",
		},
	}

	for _, tc := range tests {
		chunks, err := parseLine(tc.line, true)
		if err != nil {
			t.Fatal(err)
		}

		got := ChunksToString(chunks)
		if got != tc.want {
			t.Log(got)
			t.Log(tc.want)
			t.Fatal()
		}
	}
}

func ChunksToString(chunks []chunk) string {
	sb := &strings.Builder{}
	for _, c := range chunks {
		sb.WriteString(c.String())
	}
	return sb.String()
}

type line struct {
	txt string
	at  int
}

func (l *line) peek() (byte, bool) {
	if l.at >= len(l.txt) {
		return 0, false
	}

	n := l.txt[l.at]
	return n, true
}

func (l *line) pop() {
	l.at++
}

func (l *line) empty() bool {
	return l.at >= len(l.txt)
}

// Parses a single line, returns all the chunks (1..n) or an error.
// We stop at the first detected error.
func parseLine(txt string, complete bool) ([]chunk, error) {
	l := &line{txt, 0}
	chunks := []chunk{}
	for !l.empty() {
		c, err := parseChunk(l, complete)
		if err != nil {
			return nil, err
		}
		chunks = append(chunks, c)
	}

	return chunks, nil
}

func TestDec10(t *testing.T) {
	char2point := map[byte]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}

	test.Run(func() {
		lines, err := test.Read(test.String)
		if err != nil {
			t.Fatal(err)
		}

		score := 0
		for _, l := range lines {
			t.Log(l)
			chunks, err := parseLine(l, false)
			if err != nil {
				switch e := err.(type) {
				case unexpectedEndError:
					// ignore all lines end 'too soon'.
					t.Log(e)

				case illegalCharError:
					t.Log(e)
					score += char2point[e.got]
				}
			}

			t.Log(chunks)
		}
		test.Check(t, score, 26397, 392421)
	})
}

func ScoreDec10_2(txt string) int {
	score := 0
	for _, c := range txt {
		score *= 5

		switch c {
		case ')':
			score += 1

		case ']':
			score += 2

		case '}':
			score += 3

		case '>':
			score += 4
		}
	}
	return score
}

func TestDec10_2(t *testing.T) {
	test.Run(func() {
		lines, err := test.Read(test.String)
		if err != nil {
			t.Fatal(err)
		}

		scores := []int{}

		for _, l := range lines {
			t.Log("------------------------------------------------------------")
			t.Log(l)
			chunks, err := parseLine(l, true)
			if err != nil {
				switch e := err.(type) {
				case unexpectedEndError:
					// Should never happen due to allowing parseLine to complete the line
					// if it ends too soon.
					t.Log(e)

				case illegalCharError:
					// ignore all corrupt lines
					t.Log(err)
					continue
				}
			}

			completed := strings.TrimPrefix(ChunksToString(chunks), l)
			t.Logf("%q", completed)
			t.Logf("%q", l)
			t.Logf("%q", ChunksToString(chunks))
			if len(completed) == 0 {
				t.Log("ignoring non-completed line")
				continue
			}

			score := ScoreDec10_2(completed)
			t.Log(completed, score)
			scores = append(scores, score)
		}

		sort.Ints(scores)
		t.Log(scores)
		t.Log(len(scores))
		mid := int(len(scores) / 2)
		t.Log(mid)
		test.Check(t, scores[mid], 288957, 2769449099)
	})
}

func TestScoreDec10_2(t *testing.T) {
	tests := []struct {
		txt   string
		score int
	}{
		{
			"}}]])})]",
			288957,
		}, {
			")}>]})",
			5566,
		}, {
			"}}>}>))))",
			1480781,
		}, {
			"]]}}]}]}>",
			995444,
		}, {
			"])}>",
			294,
		},
	}

	for _, tc := range tests {
		score := ScoreDec10_2(tc.txt)
		if score != tc.score {
			t.Fatal(score, tc.score)
		}
	}
}

func TestParseLine(t *testing.T) {
	chunks, err := parseLine("((", false)
	t.Log(chunks, err)
}

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}
