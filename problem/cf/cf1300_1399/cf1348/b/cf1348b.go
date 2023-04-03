package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func solve(in io.Reader, out io.Writer) {
	var n, v, k int
	Fscan(in, &n, &k)
	var a []int
	vis := map[int]int{}
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		if vis[v] == 0 {
			a = append(a, v)
			vis[v] = 1
		}
	}
	if len(vis) > k {
		Fprintln(out, -1)
		return
	}
	for i := len(a); i < k; i++ {
		a = append(a, 1)
	}
	Fprintln(out, n*k)
	for i := 0; i < n; i++ {
		for _, v := range a {
			Fprint(out, v, " ")
		}
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
