package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func solve(in io.Reader, out io.Writer) {
	var n, m, k int
	Fscan(in, &n, &m)
	f := make([]int, m+1)
	for i := 0; i < n; i++ {
		Fscan(in, &k)
		q := make([]int, k)
		mx := make([]int, k+1)
		for j := 0; j < k; j++ {
			Fscan(in, &q[j])
			mx[k] += q[j]
		}
		for l := 0; l < k; l++ {
			p := 0
			for r := l; r < k; r++ {
				p += q[r]
				size := k - (r - l + 1)
				mx[size] = max(mx[size], mx[k]-p)
			}
		}
		for j := m; j > 0; j-- {
			for v, w := range mx {
				if j >= v {
					f[j] = max(f[j], f[j-v]+w)
				}
			}
		}
	}
	Fprintln(out, f[m])
}
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	solve(in, out)
}

func main() { run(os.Stdin, os.Stdout) }
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
