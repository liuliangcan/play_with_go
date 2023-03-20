package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// 树状数组min_right 249 ms
func solve1(in io.Reader, out io.Writer) {
	var n, v int
	Fscan(in, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		Fscan(in, &a[i])
	}
	tree := NewBIT(n*2 + 1)

	for i := 0; i < n; i++ {
		Fscan(in, &v)
		tree.AddPoint(v+1, 1)
		tree.AddPoint(v+1+n, 1)
	}

	for _, v := range a {
		p := tree.MinRight(n-v+1) - 1
		Fprint(out, (p+v)%n, " ")
		tree.AddPoint(p%n+1, -1)
		tree.AddPoint(p%n+n+1, -1)
	}
}

// 249 ms
func solve(in io.Reader, out io.Writer) {
	var n, v int
	Fscan(in, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		Fscan(in, &a[i])
	}

	dsu := NewUnionFind(n*2 + 1)
	cnt := make([]int, n*2+1)
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		cnt[v]++
		cnt[v+n]++
	}

	for _, v := range a {
		x := n - v
		for cnt[x] == 0 {
			dsu.Union(x, x+1)
			x = dsu.Find(x)
		}
		cnt[x%n]--
		cnt[x%n+n]--

		Fprint(out, (x+v)%n, " ")
	}
}
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	solve(in, out)
}

func main() { run(os.Stdin, os.Stdout) }

type BIT struct {
	C []int
	N int
}

func NewBIT(n int) BIT {
	c := make([]int, n+2)
	return BIT{c, n}
}

// AddPoint
// 给i位置增加v，注意i是1-indexed
func (u *BIT) AddPoint(i, v int) {
	for i <= u.N {
		u.C[i] += v
		i += i & -i
	}
}

// SumPrefix
// 计算[..i]的前缀和，包含i，注意i是1-indexed
func (u BIT) SumPrefix(i int) int {
	s := 0
	for i >= 1 {
		s += u.C[i]
		i &= i - 1
	}
	return s
}

// SumInterval SumIn
// 计算[l..r]闭区间的和，注意l\r是1-indexed
func (u BIT) SumInterval(l, r int) int {
	return u.SumPrefix(r) - u.SumPrefix(l-1)
}

// Lowbit
// 计算lowbit，L
func (u BIT) Lowbit(x int) int {
	return x & -x
}

// MinRight 查找[i..]后第一个>=0的位置，包含i，注意i是1-indexed
func (u BIT) MinRight(i int) int {
	p := u.SumPrefix(i)
	if i == 1 {
		if p > 0 {
			return i
		}
	} else {
		if p > u.SumPrefix(i-1) {
			return i
		}
	}
	l, r := i, u.N+1
	for l+1 < r {
		mid := (l + r) >> 1
		if u.SumPrefix(mid) > p {
			r = mid
		} else {
			l = mid
		}
	}
	return r
}

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

// Find
// 找祖宗
func (u UnionFind) Find(x int) int {
	if u.Fa[x] != x {
		u.Fa[x] = u.Find(u.Fa[x])
	}
	return u.Fa[x]
}

// Union
// newRoot = -1 表示未发生合并
func (u *UnionFind) Union(from, to int) (newRoot int) {
	x, y := u.Find(from), u.Find(to)
	if x == y {
		u.EdgeSize[y]++
		return -1
	}
	//if u.Size[x] > u.Size[y] {  // 注意如果是要定向合并，这里关闭，实际上有路径压缩，不太需要按轶合并，常关即可。
	//	x, y = y, x
	//}
	u.Fa[x] = y
	u.Size[y] += u.Size[x]
	u.EdgeSize[y] += u.EdgeSize[x] + 1
	u.Groups--
	return y
}

func (u UnionFind) Same(x, y int) bool {
	return u.Find(x) == u.Find(y)
}
