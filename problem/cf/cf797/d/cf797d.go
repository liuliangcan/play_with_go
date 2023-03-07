package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func solve(in io.Reader, out io.Writer) {
	var n, ans int
	Fscan(in, &n)
	type node struct{ v, l, r int }
	a := make([]node, n)
	cnt := map[int]int{}
	hasFa := make([]bool, n)
	for i := 0; i < n; i++ {
		Fscan(in, &a[i].v, &a[i].l, &a[i].r)
		a[i].l--
		a[i].r--
		if a[i].l >= 0 {
			hasFa[a[i].l] = true
		}
		if a[i].r >= 0 {
			hasFa[a[i].r] = true
		}
		cnt[a[i].v]++
	}
	root := 0
	for hasFa[root] {
		root++
	}
	ans = n
	var dfs func(i, l, r int)
	dfs = func(i, l, r int) {
		if i < 0 {
			return
		}
		v, x, y := a[i].v, a[i].l, a[i].r
		if l <= v && v <= r {
			ans -= cnt[v]
		}
		if l < v {
			dfs(x, l, min(r, v-1))
		}
		if r > v {
			dfs(y, max(l, v+1), r)
		}
	}
	dfs(root, 0, 1e9)
	Fprintln(out, ans)
}
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	solve(in, out)
}

func main() { run(os.Stdin, os.Stdout) }
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
