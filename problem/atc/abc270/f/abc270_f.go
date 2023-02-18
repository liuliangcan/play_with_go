package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m int
	Fscan(in, &n, &m)
	a, p := 0, n+1
	type edge struct{ u, v, w int }
	ae := make([]edge, n)
	pe := make([]edge, n)
	bridge := make([]edge, m)
	for i := 0; i < n; i++ {
		Fscan(in, &ae[i].w)
		ae[i].u = i + 1
		ae[i].v = a
	}
	for i := 0; i < n; i++ {
		Fscan(in, &pe[i].w)
		pe[i].u = i + 1
		pe[i].v = p
	}
	for i := 0; i < m; i++ {
		Fscan(in, &bridge[i].u, &bridge[i].v, &bridge[i].w)
	}
	calc := func(es []edge) int {
		uf := NewUnionFind(n + 2)
		sort.Slice(es, func(i, j int) bool { return es[i].w < es[j].w })
		res := 0
		for _, e := range es {
			u, v, w := e.u, e.v, e.w
			if uf.Merge(u, v) != -1 {
				res += w
			}
		}
		for i := 2; i <= n; i++ {
			if !uf.Same(1, i) {
				return int(1e16)
			}
		}
		return res
	}
	ans := int(1e16)
	ans = min(ans, calc(bridge))
	ans = min(ans, calc(append(bridge, ae...)))
	ans = min(ans, calc(append(bridge, pe...)))
	ans = min(ans, calc(append(append(bridge, ae...), pe...)))
	Fprintln(out, ans)

}

func main() { run(os.Stdin, os.Stdout) }

type UnionFind struct {
	Fa       []int
	Groups   int   // 连通分量个数
	Size     []int // 每组大小
	EdgeSize []int // 本家族边数(带自环/重边)
}

func NewUnionFind(n int) UnionFind {
	fa := make([]int, n) // n+1
	for i := range fa {
		fa[i] = i
	}
	size := make([]int, n)
	for i := range size {
		size[i] = 1
	}
	edgeSize := make([]int, n)
	return UnionFind{fa, n, size, edgeSize}
}

func (u UnionFind) Find(x int) int {
	if u.Fa[x] != x {
		u.Fa[x] = u.Find(u.Fa[x])
	}
	return u.Fa[x]
}

// newRoot = -1 表示未发生合并
func (u *UnionFind) Merge(from, to int) (newRoot int) {
	x, y := u.Find(from), u.Find(to)
	if x == y {
		u.EdgeSize[y]++
		return -1
	}
	if u.Size[x] > u.Size[y] {
		x, y = y, x
	}
	u.Fa[x] = y
	u.Size[y] += u.Size[x]
	u.EdgeSize[y] += u.EdgeSize[x] + 1
	u.Groups--
	return y
}

func (u UnionFind) Same(x, y int) bool {
	return u.Find(x) == u.Find(y)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
