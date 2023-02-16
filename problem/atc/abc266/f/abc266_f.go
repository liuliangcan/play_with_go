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
	var n, u, v, k int
	Fscan(in, &n)
	degree := make([]int, n)
	g := make([][]int, n)
	for i := 0; i < n; i++ {
		Fscan(in, &u, &v)
		u--
		v--
		degree[u]++
		degree[v]++
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}
	var q []int
	for i, d := range degree {
		if d == 1 {
			q = append(q, i)
		}
	}
	uf := NewUnionFind(n)
	for len(q) > 0 {
		u = q[0]
		q = q[1:]
		for _, v := range g[u] {
			uf.Merge(u, v)
			if degree[v] -= 1; degree[v] == 1 {
				q = append(q, v)
			}
		}
	}

	for Fscan(in, &k); k > 0; k-- {
		Fscan(in, &u, &v)
		if uf.Same(u-1, v-1) {
			Fprintln(out, "Yes")
		} else {
			Fprintln(out, "No")
		}
	}

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
