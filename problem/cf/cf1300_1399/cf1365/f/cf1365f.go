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
	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		Fscan(in, &a[i])
	}
	for i := 0; i < n; i++ {
		Fscan(in, &b[i])
	}
	if n%2 == 1 && a[n/2] != b[n/2] {
		Fprintln(out, "No")
		return
	}
	cnt := map[int]int{}
	l, r := 0, n-1
	for l < r {
		x, y := a[l], a[r]
		if x < y {
			x, y = y, x
		}
		cnt[(x<<32)|y]++
		x, y = b[l], b[r]
		if x < y {
			x, y = y, x
		}
		cnt[(x<<32)|y]--
		l++
		r--
	}
	for _, v := range cnt {
		if v != 0 {
			Fprintln(out, "No")
			return
		}
	}
	Fprintln(out, "yes")
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
