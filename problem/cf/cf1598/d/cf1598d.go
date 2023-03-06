package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func solve(in io.Reader, out io.Writer) {
	var n, ans int
	Fscan(in, &n)
	x := make([]int, n)
	y := make([]int, n)
	cntX := make([]int, n+1)
	cntY := make([]int, n+1)
	for i := 0; i < n; i++ {
		Fscan(in, &x[i], &y[i])
		cntX[x[i]]++
		cntY[y[i]]++
	}
	ans = n * (n - 1) * (n - 2) / 6
	for i, a := range x {
		b := y[i]
		ans -= (cntX[a] - 1) * (cntY[b] - 1)
	}

	Fprintln(out, ans)
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
