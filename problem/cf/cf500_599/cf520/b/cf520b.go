package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func solve(in io.Reader, out io.Writer) {
	var n, m, ans int
	Fscan(in, &n, &m)
	for m > n {
		if m%2 == 0 {
			m /= 2
		} else {
			m++
		}
		ans++
	}

	Fprintln(out, ans+n-m)
}
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	solve(in, out)
}

func main() { run(os.Stdin, os.Stdout) }
