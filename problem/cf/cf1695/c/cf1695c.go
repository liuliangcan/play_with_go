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
	var n, m, a, t int

	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n, &m)
		mx := make([]int, m+1)
		mn := make([]int, m+1)
		for i := 1; i <= m; i++ {
			mx[i] = -1e9
			mn[i] = 1e9
		}
		for i := 0; i < n; i++ {
			for j := 1; j <= m; j++ {
				Fscan(in, &a)
				mx[j] = max(mx[j], mx[j-1]) + a
				mn[j] = min(mn[j], mn[j-1]) + a
			}
			mx[0] = -1e9
			mn[0] = 1e9
		}
		if (m+n)&1 == 1 && mn[m] <= 0 && 0 <= mx[m] {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}

}

func main() { run(os.Stdin, os.Stdout) }
func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
