package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func solve(in io.Reader, out io.Writer) {
	var n int
	var s string
	Fscan(in, &n, &s)
	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = n - i
	}
	for i := 0; i < n; i++ {
		b[i] = i + 1
	}
	i := 0
	for i < n-1 {
		l := i
		for i < n-1 && s[i] == s[l] {
			i++
		}
		z := a
		if s[l] == '>' {
			z = b
		}
		r := i
		for l < r {
			z[l], z[r] = z[r], z[l]
			l++
			r--
		}
	}
	for _, v := range a {
		Fprint(out, v, " ")
	}
	Fprintln(out)
	for _, v := range b {
		Fprint(out, v, " ")
	}
	Fprintln(out)
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
