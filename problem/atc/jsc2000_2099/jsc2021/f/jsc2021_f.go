package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func solve(in io.Reader, out io.Writer) {
	var n, m, q, ans int
	Fscan(in, &n, &m, &q)
	ab := make([][]int, 2)
	for i := 0; i < 2; i++ {
		ab[i] = make([]int, max(n, m))
	}
	type query struct{ t, i, y int }
	qs := make([]query, q)
	mp := map[int]int{0: 1}
	for i := 0; i < q; i++ {
		Fscan(in, &qs[i].t, &qs[i].i, &qs[i].y)
		mp[qs[i].y] = 1
	}
	var hs []int
	for v := range mp {
		hs = append(hs, v)
	}
	sort.Ints(hs)
	sz := len(hs)
	trees := make([][]BIT, 2)
	for i := 0; i < 2; i++ {
		trees[i] = []BIT{NewBIT(sz), NewBIT(sz)}
	}
	trees[0][0].AddPoint(1, n)
	trees[1][0].AddPoint(1, m)
	for _, v := range qs {
		t, i, y := v.t, v.i, v.y
		a := ab[t-1]
		x := a[i-1]
		a[i-1] = y
		xx := sort.Search(len(hs), func(i int) bool {
			return hs[i] >= x
		}) + 1
		yy := sort.Search(len(hs), func(i int) bool {
			return hs[i] >= y
		}) + 1

		cnt, sum := trees[t-1][0], trees[t-1][1]
		cnt.AddPoint(xx, -1)
		cnt.AddPoint(yy, +1)
		sum.AddPoint(xx, -x)
		sum.AddPoint(yy, +y)

		cnt, sum = trees[2-t][0], trees[2-t][1]
		ans -= cnt.SumPrefix(xx)*x + sum.SumInterval(xx+1, sz)
		ans += cnt.SumPrefix(yy)*y + sum.SumInterval(yy+1, sz)
		Fprintln(out, ans)
	}
}
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	solve(in, out)
}

func main() { run(os.Stdin, os.Stdout) }
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type BIT struct {
	C []int
	N int
}

func NewBIT(n int) BIT {
	c := make([]int, n+2)
	return BIT{c, n}
}

// AddPoint
// 给i位置增加v，注意i是1-indexed
func (u *BIT) AddPoint(i, v int) {
	for i <= u.N {
		u.C[i] += v
		i += i & -i
	}
}

// SumPrefix
// 计算[..i]的前缀和，包含i，注意i是1-indexed
func (u BIT) SumPrefix(i int) int {
	s := 0
	for i >= 1 {
		s += u.C[i]
		i &= i - 1
	}
	return s
}

// SumInterval SumIn
// 计算[l..r]闭区间的和，注意l\r是1-indexed
func (u BIT) SumInterval(l, r int) int {
	return u.SumPrefix(r) - u.SumPrefix(l-1)
}

// Lowbit
// 计算lowbit，L
func (u BIT) Lowbit(x int) int {
	return x & -x
}
