package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func solve(in io.Reader, out io.Writer) {
	var n, d, h int
	Fscan(in, &n, &d, &h)
	if 2*h < d || d == 1 && n > 2 {
		Fprintln(out, -1)
		return
	}
	if n == 2 {
		Fprintln(out, "1 2")
		return
	}
	if h == 1 {
		for i := 2; i <= n; i++ {
			Fprintln(out, "1", i)
		}
		return
	}
	for i := 1; i <= h; i++ {
		Fprintln(out, i, i+1)
	}
	if d > h {
		Fprintln(out, 1, h+2)
		for i := h + 2; i <= d; i++ {
			Fprintln(out, i, i+1)
		}
	}
	for i := d + 2; i <= n; i++ {
		Fprintln(out, 2, i)
	}
}
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	solve(in, out)
}

func main() { run(os.Stdin, os.Stdout) }
