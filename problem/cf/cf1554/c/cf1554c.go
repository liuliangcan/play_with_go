package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func solve(in io.Reader, out io.Writer) {
	var n, m, mex int
	Fscan(in, &n, &m)
	m++
	for i := 29; i >= 0; i-- {
		x, y := n>>i&1, m>>i&1
		if x > 0 && y == 0 {
			break
		}
		if x == 0 && y == 1 {
			mex |= 1 << i
		}
	}
	Fprintln(out, mex)
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
