package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"os"
)

func solve(in io.Reader, out io.Writer) {
	var n, m, u, v, w int
	Fscan(in, &n, &m)
	type ee struct{ v, w int }
	g := make([][]ee, n)
	for i := 0; i < m; i++ {
		Fscan(in, &u, &v, &w)
		u--
		v--
		g[u] = append(g[u], ee{v, w})
		g[v] = append(g[v], ee{u, w})
	}
	t := make([]map[int]int, n)
	for i := 0; i < n; i++ {
		Fscan(in, &w)
		t[i] = map[int]int{}
		for j := 0; j < w; j++ {
			Fscan(in, &u)
			t[i][u] = 1
		}
	}
	dist := make([]int, n)
	for i := 1; i < n; i++ {
		dist[i] = int(1e13)
	}
	h := hp{{0, 0}}
	for len(h) > 0 {
		top := heap.Pop(&h).(pair)
		d, u := top.dis, top.v
		if d > dist[u] {
			continue
		}
		if u == n-1 {
			Fprintln(out, d)
			return
		}
		for t[u][d] > 0 {
			d++
		}
		for _, e := range g[u] {
			v, w = e.v, e.w
			if d+w < dist[v] {
				dist[v] = d + w
				heap.Push(&h, pair{v, d + w})
			}
		}
	}

	Fprintln(out, -1)
}
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	solve(in, out)
}

func main() { run(os.Stdin, os.Stdout) }

type pair struct{ v, dis int }
type hp []pair

func (h hp) Len() int              { return len(h) }
func (h hp) Less(i, j int) bool    { return h[i].dis < h[j].dis }
func (h hp) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{})   { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
