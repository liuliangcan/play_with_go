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
	var n, ans, u, v, cnt_u, cnt_e int
	Fscan(in, &n)
	g := make([][]int, 4e5+1)
	for i := 0; i < n; i++ {
		Fscan(in, &u, &v)
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}
	vis := make([]bool, len(g))
	var dfs func(int)
	dfs = func(u int) {
		vis[u] = true
		cnt_u += 1
		cnt_e += len(g[u])
		for _, v := range g[u] {
			if !vis[v] {
				dfs(v)
			}
		}
	}

	// dfs方法计算每个连通块的点数和边数，取小的那个(无环树取边数；有环取点数)
	// 并查集方法用py实现；dfs方法用go实现；bfs方法用rust实现；可移步查看
	for u, b := range vis {
		if !b && g[u] != nil {
			cnt_u, cnt_e = 0, 0
			dfs(u)
			ans += min(cnt_u, cnt_e/2) // 树则边数，有环则节点数，因此边不用统计完，因此可以用vis
		}
	}
	Fprintln(out, ans)

}

func main() { run(os.Stdin, os.Stdout) }
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
