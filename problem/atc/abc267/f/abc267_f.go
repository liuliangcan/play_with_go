package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, u, v, k int

	Fscan(in, &n)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		Fscan(in, &u, &v)
		v--
		u--
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}
	Fscan(in, &q)
	type pair struct{ i, k int }
	qs := make([][]pair, n)
	for i := 0; i < q; i++ {
		Fscan(in, &u, &k)
		u--
		qs[u] = append(qs[u], pair{i, k})
	}
	ans := make([]int, q)
	for i, _ := range ans {
		ans[i] = -1
	}
	path := make([]int, n)
	leaf := 0
	mx := 0
	for i := 0; i < 3; i++ {
		var dfs func(int, int, int)
		dfs = func(u, fa, d int) {
			path[d] = u
			if d > mx {
				mx = d
				leaf = u
			}
			for _, p := range qs[u] {
				if d >= p.k {
					ans[p.i] = path[d-p.k] + 1
				}
			}
			for _, v := range g[u] {
				if v != fa {
					dfs(v, u, d+1)
				}
			}
		}
		dfs(leaf, -1, 0)
	}
	for _, v := range ans {
		Fprintln(out, v)
	}
}

func main() { run(os.Stdin, os.Stdout) }
