package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// SortedList 1372ms
func solve1(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	type seg struct{ i, l, r int }
	a := make([]seg, n)
	for i := 0; i < n; i++ {
		Fscan(in, &a[i].l, &a[i].r)
		a[i].i = i
	}
	sort.Slice(a, func(i, j int) bool { return a[i].r < a[j].r })
	ans := make([]int, n)
	p := NewSortedList()
	for _, s := range a {
		l, i := s.l, s.i
		//ans[i] = len(p) - p.BisectLeft(l)
		ans[i] = p.Len() - p.BisectLeft(l)
		p.Add(l)
	}

	for _, v := range ans {
		Fprintln(out, v)
	}
}

// BIT 420 ms
func solve(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	type seg struct{ i, l, r int }
	var hs []int
	a := make([]seg, n)
	for i := 0; i < n; i++ {
		Fscan(in, &a[i].l, &a[i].r)
		hs = append(hs, []int{a[i].l, a[i].r}...)
		a[i].i = i
	}
	sort.Slice(a, func(i, j int) bool { return a[i].r < a[j].r })
	sort.Ints(hs)
	ans := make([]int, n)
	p := NewBIT(n * 2)
	for _, s := range a {
		l, i := s.l, s.i
		l = sort.Search(len(hs), func(i int) bool { return hs[i] >= l })
		ans[i] = p.SumInterval(l+1, n*2+1)
		p.AddPoint(l+1, 1)
	}

	for _, v := range ans {
		Fprintln(out, v)
	}
}
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	solve(in, out)
}

func main() { run(os.Stdin, os.Stdout) }

type SortedList struct {
	Small []int
	Large []int
}

func NewSortedList() SortedList {
	var small []int
	var large []int
	return SortedList{small, large}
}

func (u *SortedList) Add(v int) {
	if len(u.Small) > 6700 {
		u.Large = append(u.Large, u.Small...)
		u.Small = []int{}
		sort.Ints(u.Large)
	}
	if len(u.Small) == 0 || u.Small[len(u.Small)-1] <= v {
		u.Small = append(u.Small, v)
	} else if u.Small[0] >= v {
		u.Small = append([]int{v}, u.Small...)
	} else {
		p := sort.Search(len(u.Small), func(i int) bool { return u.Small[i] >= v })
		t := append([]int{}, u.Small[p:]...) // 实现insert 需要保存临时变量
		u.Small = append(u.Small[:p], v)
		u.Small = append(u.Small, t...)
		//u.Small = append(u.Small[:p], append([]int{v}, u.Small[p:]...)...)
	}
}

func (u SortedList) BisectLeft(v int) int {
	return sort.Search(len(u.Small), func(i int) bool { return u.Small[i] >= v }) + sort.Search(len(u.Large), func(i int) bool { return u.Large[i] >= v })
}

func (u SortedList) Len() int {
	return len(u.Small) + len(u.Large)
}

type BIT struct {
	C []int
	N int
}

func NewBIT(n int) BIT {
	c := make([]int, n+2)
	return BIT{c, n}
}

func (u *BIT) AddPoint(i, v int) {
	for i <= u.N {
		u.C[i] += v
		i += i & -i
	}
}

func (u BIT) SumPrefix(i int) int {
	s := 0
	for i >= 1 {
		s += u.C[i]
		i &= i - 1
	}
	return s
}
func (u BIT) SumInterval(l, r int) int {
	return u.SumPrefix(r) - u.SumPrefix(l-1)
}

func (u BIT) Lowbit(x int) int {
	return x & -x
}
