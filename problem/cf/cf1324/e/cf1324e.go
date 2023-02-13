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
	var n, h, l, r int
	Fscan(in, &n, &h, &l, &r)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		Fscan(in, &a[i])
	}
	f := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		f[i] = make([]int, h)
	}
	for s := l; s <= r; s++ {
		f[n][s] = 1
	}
	for i := n - 1; i >= 0; i-- {
		for s := 0; s < h; s++ {
			f[i][s] = max(f[i+1][(s+a[i])%h], f[i+1][(s+a[i]-1+h)%h])
			if i > 0 && l <= s && s <= r {
				f[i][s] += 1
			}
		}
	}

	Fprintln(out, f[0][0])
}

func main() { run(os.Stdin, os.Stdout) }
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
