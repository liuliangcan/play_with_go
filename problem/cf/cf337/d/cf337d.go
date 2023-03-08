package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func solve(in io.Reader, out io.Writer) {
	var n, m, d, u, v int
	Fscan(in, &n, &m, &d)
	g := make([][]int, n)
	up := make([]int, n)
	down1 := make([]int, n)
	down2 := make([]int, n)
	for i := 0; i < n; i++ {
		down1[i] = -1e9
		down2[i] = -1e9
		up[i] = -1e9
	}
	for i := 0; i < m; i++ {
		Fscan(in, &u)
		u--
		down1[u] = 0
	}
	for i := 0; i < n-1; i++ {
		Fscan(in, &u, &v)
		u--
		v--
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}
	var dfs func(u, fa int)
	dfs = func(u, fa int) {
		for _, v := range g[u] {
			if v == fa {
				continue
			}
			dfs(v, u)
			c := down1[v] + 1
			if c > down1[u] {
				down2[u], down1[u] = down1[u], c
			} else if c > down2[u] {
				down2[u] = c
			}
		}
	}
	var reroot func(u, fa int)
	reroot = func(u, fa int) {
		for _, v := range g[u] {
			if v == fa {
				continue
			}
			if down1[v]+1 == down1[u] {
				up[v] = max(up[u], down2[u]) + 1
			} else {
				up[v] = max(up[u], down1[u]) + 1
			}
			reroot(v, u)
		}
	}
	dfs(0, -1)
	reroot(0, -1)
	ans := 0
	for i := 0; i < n; i++ {
		if max(up[i], down1[i]) <= d {
			ans++
		}
	}
	Fprintln(out, ans)
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
