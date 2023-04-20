package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func solve(in io.Reader, out io.Writer) {
	var n, a, b, da, db, u, v int
	Fscan(in, &n, &a, &b, &da, &db)
	a--
	b--
	g := make([][]int, n)
	for i := 0; i < n-1; i++ {
		Fscan(in, &u, &v)
		u--
		v--
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}
	if da*2 >= db {
		Fprintln(out, "Alice")
		return
	}

	get_d := func() int {
		var d int
		vis := make([]int, n)
		vis[a] = 1
		q := []int{a}
		for len(q) > 0 {
			d++
			var nq []int
			for _, u := range q {
				for _, v := range g[u] {
					if v == b {
						return d
					}
					if vis[v] == 0 {
						vis[v] = 1
						nq = append(nq, v)
					}
				}
			}
			q = nq
		}
		return d
	}
	d := get_d()
	if d <= da {
		Fprintln(out, "Alice")
		return
	}
	start := 0
	dim := 0
	for i := 0; i < 2; i++ {
		dim = -1
		vis := make([]int, n)
		vis[start] = 1
		q := []int{start}
		for len(q) > 0 {
			dim++
			var nq []int
			for _, u := range q {
				for _, v := range g[u] {
					if vis[v] == 0 {
						vis[v] = 1
						start = v
						nq = append(nq, v)
					}
				}
			}
			q = nq
		}
	}
	if dim <= 2*da {
		Fprintln(out, "Alice")
		return

	}
	Fprintln(out, "Bob")
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
