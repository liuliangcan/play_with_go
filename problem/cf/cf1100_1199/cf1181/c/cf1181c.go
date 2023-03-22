package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func solve(in io.Reader, out io.Writer) {
	var n, m, ans int
	Fscan(in, &n, &m)
	g := make([]string, n)
	left := make([]int, n)
	for i := 0; i < n; i++ {
		Fscan(in, &g[i])
		left[i] = 1
	}
	for j := 0; j < m; j++ {
		up, f1, f2 := make([]int, n), make([]int, n), make([]int, n)
		up[0] = 1
		for i := 0; i < n; i++ {
			if j > 0 && g[i][j] == g[i][j-1] {
				left[i]++
			} else {
				left[i] = 1
			}
			f1[i], f2[i] = left[i], left[i]
		}

		for i := n - 2; i >= 0; i-- {
			if g[i][j] == g[i+1][j] {
				f2[i] = min(f2[i], f2[i+1])
			}
		}
		for i := 1; i < n; i++ {
			if g[i][j] == g[i-1][j] {
				up[i] = up[i-1] + 1
				f1[i] = min(f1[i], f1[i-1])
			} else {
				up[i] = 1
			}
			u := up[i]
			if i+1 < u*3 || up[i-u] != u || up[i-u*2] < u {
				continue
			}
			ans += min(f1[i], min(f1[i-u], f2[i-3*u+1]))
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
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
