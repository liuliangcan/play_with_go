package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func solve(in io.Reader, out io.Writer) {
	var n, c int
	Fscan(in, &n, &c)
	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n-1; i++ {
		Fscan(in, &a[i])
	}
	for i := 0; i < n-1; i++ {
		Fscan(in, &b[i])
	}
	Fprint(out, 0, " ")
	ladder, ele := 0, c
	for i := 1; i < n; i++ {
		ladder, ele = min(ladder, ele)+a[i-1], min(ele, ladder+c)+b[i-1]
		Fprint(out, min(ladder, ele), " ")
	}
}
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	solve(in, out)
}

func main() { run(os.Stdin, os.Stdout) }
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
