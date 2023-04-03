package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func solve(in io.Reader, out io.Writer) {
	var n, v, p int
	Fscan(in, &n)
	var a []int
	for i := 0; i < n; i++ {
		p = v
		Fscan(in, &v)
		if i > 0 {
			a = append(a, abs(v-p))
		}
	}
	f := func(start int) int {
		mx, dp := 0, 0
		sign := 1
		for i := start; i < n-1; i++ {
			dp = a[i]*sign + max(0, dp)
			mx = max(mx, dp)
			sign *= -1
		}
		return mx
	}

	Fprintln(out, max(f(0), f(1)))
}
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

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
