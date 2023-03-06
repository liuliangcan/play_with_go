package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func solve(in io.Reader, out io.Writer) {
	var n, x, y int
	var s, t string
	Fscan(in, &n, &x, &y, &s, &t)
	var diff []int
	for i := range s {
		if s[i] != t[i] {
			diff = append(diff, i)
		}
	}
	m := len(diff)
	if m%2 > 0 {
		Fprintln(out, "-1")
		return
	}
	if y <= x {
		if m == 2 && diff[0]+1 == diff[1] {
			Fprintln(out, min(x, 2*y))
			return
		}
		Fprintln(out, m/2*y)
	} else {
		f := make([]int, m+1)
		for i, v := range diff {
			if i == 0 {
				f[1] = y
			} else {
				f[i+1] = min(f[i]+y, f[i-1]+(v-diff[i-1])*2*x)
			}
		}
		Fprintln(out, f[m]/2)
	}
}
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var t int
	for Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
	//solve(in, out)
}

func main() { run(os.Stdin, os.Stdout) }
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
