package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func solve(in io.Reader, out io.Writer) {
	var n, k, fa int
	Fscan(in, &n, &k)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		Fscan(in, &fa)
		g[fa-1] = append(g[fa-1], i)
	}
	ok := func(x int) bool {
		if x == 0 {
			return n == 1
		}
		cnt := 0
		var dfs func(u int) int
		dfs = func(u int) int {
			if cnt > k {
				return 0
			}
			h := 0
			for _, v := range g[u] {
				p := dfs(v)
				if p == x && u > 0 {
					cnt++
				} else {
					h = max(h, p)
				}
			}
			return h + 1
		}
		dfs(0)
		return cnt <= k
	}
	ans := sort.Search(n, ok)
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
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
