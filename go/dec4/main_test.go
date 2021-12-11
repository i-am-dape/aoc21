package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/i-am-dape/aoc21/go/test"
)

type board struct {
	n    int
	rows [][]element
}

type element struct {
	val    int
	marked bool
}

func (b *board) mark(n int) {
	for r := 0; r < len(b.rows); r++ {
		for c := 0; c < len(b.rows[r]); c++ {
			e := b.rows[r][c]
			if n == e.val {
				e.marked = true
				b.rows[r][c] = e
				return
			}
		}
	}
}

func (b *board) check() bool {
	return b.checkRows() || b.checkColumns()
}

func (b *board) checkColumns() bool {
	if len(b.rows) == 0 {
		return false
	}

	cols := len(b.rows[0])
	for c := 0; c < cols; c++ {
		marked := 0
		for r := 0; r < len(b.rows); r++ {
			if !b.rows[r][c].marked {
				break
			}

			marked++
		}

		if marked == len(b.rows) {
			return true
		}
	}

	return false
}

func (b *board) checkRows() bool {
	for r := 0; r < len(b.rows); r++ {
		marked := 0
		checked := 0
		for c := 0; c < len(b.rows[r]); c++ {
			checked++
			if !b.rows[r][c].marked {
				break
			}

			marked++
		}

		if checked > 0 && marked == len(b.rows[r]) {
			return true
		}
	}

	return false
}

func TestMarkAndCheck(t *testing.T) {
	tests := []struct {
		drawn  []int
		winner bool
		usum   int
	}{
		{[]int{1}, false, 44},
		{[]int{1, 2, 3}, true, 39},
		{[]int{4, 5, 6}, true, 30},
		{[]int{7, 8, 9}, true, 21},
		{[]int{1, 5, 7, 8, 9}, true, 15},
		{[]int{1, 4, 7}, true, 33},
		{[]int{2, 5, 8}, true, 30},
		{[]int{3, 6, 9}, true, 27},
		{[]int{7, 5, 3, 6, 9}, true, 15},
	}

	for _, tc := range tests {
		t.Log(tc)

		b := &board{0, [][]element{
			{{1, false}, {2, false}, {3, false}},
			{{4, false}, {5, false}, {6, false}},
			{{7, false}, {8, false}, {9, false}},
		}}

		for _, d := range tc.drawn {
			b.mark(d)
		}

		winner := b.check()
		if winner != tc.winner {
			t.Fatal(tc, winner, tc.winner)
		}

		usum := b.usum()
		if usum != tc.usum {
			t.Fatal(usum, tc.usum)
		}
	}
}

func (b *board) usum() int {
	usum := 0
	for _, row := range b.rows {
		for _, elem := range row {
			if !elem.marked {
				usum += elem.val
			}
		}
	}
	return usum
}

type Queue[T any] struct {
	items []T
	at    int
}

func (q *Queue[T]) Size() int {
	return len(q.items) - q.at
}

func (q *Queue[T]) Pop() (T, bool) {
	var v T
	if q.at == len(q.items) {
		return v, false
	}

	v = q.items[q.at]
	q.at++
	return v, true
}

type BingoInput struct {
	at     int
	drawn  Queue[int]
	boards []*board
}

func (bi *BingoInput) winners() ([]*board, int) {
	for {
		d, ok := bi.drawn.Pop()
		if !ok {
			return nil, -1
		}

		for _, b := range bi.boards {
			b.mark(d)
		}

		winners := []*board{}
		rem := []*board{}
		for _, b := range bi.boards {
			if b.check() {
				winners = append(winners, b)
			} else {
				rem = append(rem, b)
			}
		}

		bi.boards = rem

		if len(winners) > 0 {
			return winners, d
		}
	}
}

func removeBoard(boards []*board, at int) []*board {
	head := boards[:at]
	tail := boards[at+1:]
	return append(head, tail...)
}

func TestRemoveBoard(t *testing.T) {
	log := func(bs []*board) {
		for _, b := range bs {
			t.Log(b.n)
		}
	}

	b := []*board{{0, nil}, {1, nil}, {2, nil}}
	b = removeBoard(b, 0)
	t.Log(b)
	log(b)
	if len(b) != 2 || b[0].n != 1 || b[1].n != 2 {
		t.Fatal(b)
	}

	b = []*board{{0, nil}, {1, nil}, {2, nil}}
	b = removeBoard(b, 1)
	t.Log(b)
	log(b)
	if len(b) != 2 || b[0].n != 0 || b[1].n != 2 {
		t.Fatal(b)
	}

	b = []*board{{0, nil}, {1, nil}, {2, nil}}
	b = removeBoard(b, 2)
	t.Log(b)
	log(b)
	if len(b) != 2 || b[0].n != 0 || b[1].n != 1 {
		t.Fatal(b)
	}
}

func ReadBingoInput(t *testing.T) *BingoInput {
	parseInt := func(txt string) int {
		n, err := strconv.ParseInt(txt, 10, 32)
		if err != nil {
			t.Fatal(err)
		}
		return int(n)
	}

	txt, err := test.Read(test.String)
	if err != nil {
		t.Fatal(err)
	}

	drawn := []int{}
	for _, ntxt := range strings.Split(txt[0], ",") {
		drawn = append(drawn, parseInt(ntxt))
	}

	bi := &BingoInput{0, Queue[int]{drawn, 0}, []*board{}}

	if len(txt[1]) != 0 {
		t.Fatal("this line should be empty; line #", 1, txt)
	}

	bn := 0
	for i := 2; i < len(txt); i++ {
		// assuming well formed input
		cur := &board{bn, [][]element{}}
		for j := 0; j < 5; j++ {
			row := []element{}
			for _, ntxt := range strings.Split(txt[i+j], " ") {
				// skip if this is a single digit, this is indicated
				// by ntxt being empty
				if len(ntxt) == 0 {
					continue
				}

				row = append(row, element{parseInt(ntxt), false})
			}
			cur.rows = append(cur.rows, row)
		}
		bi.boards = append(bi.boards, cur)
		bn++
		i += 5
	}

	return bi
}

func TestStringsSplit(t *testing.T) {
	parts := strings.Split("22 13 17 11  0", " ")
	t.Log(parts)
}

func TestDec4(t *testing.T) {
	test.Run(func() {
		bi := ReadBingoInput(t)
		ws, d := bi.winners()
		for _, w := range ws {
			t.Log(w, ws[0].usum()*d)
		}

		test.Check(t, ws[0].usum()*d, 4512, 72770)
	})
}

func TestDec4_2(t *testing.T) {
	test.Run(func() {
		bi := ReadBingoInput(t)
		fmt.Println("initial drawn:", bi.drawn)
		winners := [][]*board{}
		winnerDs := []int{}

		for {
			t.Log("rem:", len(bi.boards), bi.drawn.Size())
			ws, d := bi.winners()
			if ws != nil {
				winners = append(winners, ws)
				winnerDs = append(winnerDs, d)
				continue
			}
			last := len(winners) - 1
			lw := winners[last]
			ld := winnerDs[last]
			t.Log(lw, ld, lw[0].usum()*ld)
			test.Check(t, lw[0].usum()*ld, 1924, 13912)

			for _, b := range bi.boards {
				if b.check() {
					t.Error(b)
				}
			}

			return
		}
	})
}

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}
