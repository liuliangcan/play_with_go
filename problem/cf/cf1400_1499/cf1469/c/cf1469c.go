package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func solve(in io.Reader, out io.Writer) {
	var n, k int
	Fscan(in, &n, &k)
	h := make([]int, n)
	for i := 0; i < n; i++ {
		Fscan(in, &h[i])
	}
	mx, mn := h[0], h[0]
	for i := 1; i < n; i++ {
		mn = max(h[i], mn-k+1)
		mx = min(h[i]+k-1, mx+k-1)
		if mn > mx {
			Fprintln(out, "NO")
			return
		}
	}
	if mn <= h[len(h)-1] && h[len(h)-1] <= mx {
		Fprintln(out, "YES")
		return
	}

	Fprintln(out, "NO")

}
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var t int
	for Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}

func main() { run(os.Stdin, os.Stdout) }
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
