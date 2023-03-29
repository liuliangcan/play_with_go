package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func solve(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	type pairs struct{ p, c int }
	a := make([]pairs, n)
	for i := 0; i < n; i++ {
		Fscan(in, &a[i].p, &a[i].c)
	}
	sort.Slice(a, func(i, j int) bool { return a[i].p < a[j].p })
	f := make([]int, n) // f[i] 固定第i个球时前边的最小花费
	f[0] = a[0].c
	g := int(1e15) // 不固定时的花费
	for i := 1; i < n; i++ {
		f[i] = a[i].c + min(f[i-1], g)
		g = int(1e15)
		s := a[i].p
		for j := i - 1; j >= 0; j-- {
			g = min(g, s-(i-j)*a[j].p+f[j])
			s += a[j].p
		}
	}

	Fprintln(out, min(f[n-1], g))
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
