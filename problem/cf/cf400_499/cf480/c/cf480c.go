package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

const MOD = int(1e9 + 7)

func solve(in io.Reader, out io.Writer) {
	var n, a, b, k int
	Fscan(in, &n, &a, &b, &k)
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = 1
	}
	f[b] = 0
	for i := 0; i < k; i++ {
		p := make([]int, n+2)
		for j, v := range f {
			p[j+1] = p[j] + v
		}
		for j := 1; j <= n; j++ {
			d := abs(j - b)
			l, r := max(1, j-d+1), min(n, j+d-1)
			f[j] = (p[r+1] - p[j+1] + p[j] - p[l]) % MOD
		}
	}

	Fprintln(out, f[a])
}
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	//var t int
	//for Fscan(in, &t); t > 0; t-- {
	//	solve(in, out)
	//}
	solve(in, out)
}

func main() { run(os.Stdin, os.Stdout) }
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
