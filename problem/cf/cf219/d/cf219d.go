package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func solve(in io.Reader, out io.Writer) {
	var n, u, v int
	Fscan(in, &n)
	type edge struct{ v, dir int } // 邻居，方向(1正边，-1反边)
	g := make([][]edge, n)
	for i := 0; i < n-1; i++ {
		Fscan(in, &u, &v)
		u--
		v--
		g[u] = append(g[u], edge{v, 1})
		g[v] = append(g[v], edge{u, -1})
	}
	f := make([]int, n) // 反边
	var dfs func(u, fa int)
	dfs = func(u, fa int) {
		for _, e := range g[u] {
			v, d := e.v, e.dir
			if v != fa {
				dfs(v, u)
				//f[u] += f[v]
				//if d < 0 {
				//	f[u]++
				//}
				if d < 0 {
					f[0]++
				}
			}
		}
	}
	var reroot func(u, fa int)
	reroot = func(u, fa int) {
		for _, e := range g[u] {
			v, d := e.v, e.dir
			if v != fa {
				f[v] = f[u] + d
				reroot(v, u)
			}
		}
	}
	dfs(0, -1)
	reroot(0, -1)
	var ans []int
	mn := f[0]
	for i, v := range f {
		if v < mn {
			mn = v
			ans = []int{i + 1}
		} else if v == mn {
			ans = append(ans, i+1)
		}
	}

	Fprintln(out, mn)
	for _, v := range ans {
		Fprint(out, v, " ")
	}
}
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	solve(in, out)
}

func main() { run(os.Stdin, os.Stdout) }
