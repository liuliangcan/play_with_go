package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func solve(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	c := make([]int, n)
	s := 0
	for i := 0; i < n; i++ {
		Fscan(in, &c[i])
		s += c[i]
	}
	k := s / n
	d := make([]int, n)
	sd := 0
	a := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		sd += d[i]
		if c[i]+sd == i+1 {
			a[i] = 1
		}
		sd--
		if i-k >= 0 {
			d[i-k]++
		}
		if a[i] == 1 {
			k--
		}
	}
	for _, v := range a {
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
