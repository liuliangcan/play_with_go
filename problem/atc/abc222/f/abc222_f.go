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
	var n, u, v, w int
	Fscan(in, &n)
	type nb struct{ v, w int }
	g := make([][]nb, n)
	for i := 1; i < n; i++ {
		Fscan(in, &u, &v, &w)
		u--
		v--
		g[u] = append(g[u], nb{v, w})
		g[v] = append(g[v], nb{u, w})
	}
	d := make([]int, n)
	for i := 0; i < n; i++ {
		Fscan(in, &d[i])
	}
	down1 := make([]int, n)
	down2 := make([]int, n)
	up := make([]int, n)
	var dfs func(u int, fa int)
	dfs = func(u int, fa int) {
		for _, o := range g[u] {
			v, w := o.v, o.w
			if v == fa {
				continue
			}
			dfs(v, u)
			s := max(down1[v], d[v]) + w
			if s > down1[u] {
				down2[u] = down1[u]
				down1[u] = s
			} else if s > down2[u] {
				down2[u] = s
			}
		}
	}
	dfs(0, -1)
	var reroot func(u int, fa int)
	reroot = func(u int, fa int) {
		for _, o := range g[u] {
			v, w := o.v, o.w
			if v == fa {
				continue
			}
			if down1[u] == max(down1[v], d[v])+w {
				// v 在 u 的最大路径上， 则v向上的最大路径产生在up[u]和down2[u]之间,别忘了d[u]
				up[v] = max(max(up[u], down2[u]), d[u]) + w
			} else {
				// v 不在 u 的最大路径上， 则v向上的最大路径产生在up[u]和down1[u]之间,别忘了d[u]
				up[v] = max(max(up[u], down1[u]), d[u]) + w
			}
			reroot(v, u)
		}
	}
	reroot(0, -1)
	for i := 0; i < n; i++ {
		Fprintln(out, max(up[i], down1[i]), " ")
	}
}

func main() { run(os.Stdin, os.Stdout) }
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
