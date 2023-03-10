package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func solve(in io.Reader, out io.Writer) {
	var n, m int
	Fscan(in, &n, &m)
	a := make([]int, m)
	type pair struct{ x, y int }
	ba := make([]pair, m)
	for i := 0; i < m; i++ {
		Fscan(in, &ba[i].x, &ba[i].y)
		a[i] = ba[i].x
	}
	sort.Ints(a)
	sort.Slice(ba, func(i, j int) bool { return ba[i].y > ba[j].y || ba[i].y == ba[j].y && ba[i].x > ba[j].x })

	ans := int64(0)
	i := m - 1
	s := int64(0)
	for _, pr := range ba {
		x, y := pr.x, pr.y
		for n > 0 && i >= 0 && a[i] >= y {
			s += int64(a[i])
			i--
			n--
		}
		if n <= 0 {
			ans = max64(ans, s)
			break
		}
		if x >= y {
			ans = max64(ans, s+int64(n)*int64(y))
		} else {
			ans = max64(ans, s+int64(x)+int64(n-1)*int64(y))
		}
	}

	Fprintln(out, ans)
}

// 二分140ms
func solve1(in io.Reader, out io.Writer) {
	var n, m int
	Fscan(in, &n, &m)
	a := make([]int, m)
	type pair struct{ x, y int }
	ba := make([]pair, m)
	for i := 0; i < m; i++ {
		Fscan(in, &ba[i].x, &ba[i].y)
		a[i] = ba[i].x
	}
	sort.Ints(a)
	pa := make([]int64, m+1)
	for i := m - 1; i >= 0; i-- {
		pa[i] = pa[i+1] + int64(a[i])
	}
	ans := int64(0)
	if n <= m {
		ans = pa[m-n]
	}
	for _, pr := range ba {
		x, y := pr.x, pr.y
		p := sort.Search(m, func(i int) bool {
			return a[i] >= y
		})
		z := m - p
		if z >= n {
			continue
		}
		if x >= y {
			ans = max64(ans, pa[m-z]+int64(n-z)*int64(y))
		} else {
			ans = max64(ans, pa[m-z]+int64(x)+int64(n-z-1)*int64(y))
		}
	}

	Fprintln(out, ans)
}
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var t int
	for Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}

func main() { run(os.Stdin, os.Stdout) }
func max64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
