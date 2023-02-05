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

	var n, m, v int
	Fscan(in, &n, &m)
	b := make([][]int, m+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &v)
		l := 1
		if v < 0 {
			l = (-v + i - 1) / i
		}
		for ; l <= m && v+i*l < n; l++ {
			b[l] = append(b[l], v+i*l)
		}
	}
	time := make([]int, n+1)
	for i := 1; i <= m; i++ {
		for _, v := range b[i] {
			time[v] = i
		}
		mex := 0
		for time[mex] == i {
			mex++
		}
		Fprintln(out, mex)
	}

}

func main() { run(os.Stdin, os.Stdout) }
