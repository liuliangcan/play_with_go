package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func solve(in io.Reader, out io.Writer) {
	var n, w, m int
	Fscan(in, &n, &m, &w)
	a := make([]int, n)
	mx := 1
	for i := 0; i < n; i++ {
		Fscan(in, &a[i])
		if a[i] > mx {
			mx = a[i]
		}
	}

	Fprintln(out, sort.Search(mx+m+1, func(x int) bool {
		cnt, s := 0, 0
		d := make([]int, n+w+1)
		for i, v := range a {
			s += d[i]
			delta := x - s - v
			if delta > 0 {
				cnt += delta
				if cnt > m {
					return true
				}
				s += delta
				d[i+w] -= delta
			}
		}
		return false
	})-1)
}
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	solve(in, out)
}

func main() { run(os.Stdin, os.Stdout) }
