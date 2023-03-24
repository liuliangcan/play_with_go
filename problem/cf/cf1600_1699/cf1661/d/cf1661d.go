package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func solve(in io.Reader, out io.Writer) {
	var n, k, ans, d, a int
	Fscan(in, &n, &k)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		Fscan(in, &b[i])
	}
	d2 := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		d += d2[i]
		a += d
		if a < b[i] {
			k2 := min(i+1, k)
			times := (b[i] - a + k2 - 1) / k2
			ans += times
			a += times * k2
			if i > 0 {
				d2[i-1] -= times
				if i > k2 {
					d2[i-k2-1] += times
				}
			}
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
