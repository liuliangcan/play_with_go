package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func solve(in io.Reader, out io.Writer) {
	var n, m, v, s int
	Fscan(in, &n, &m)
	z := []int{0}
	y := []int{0}
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		s += v
		if s > z[len(z)-1] {
			z = append(z, s)
			y = append(y, i)
		}

	}
	mx := z[len(z)-1]
	for i := 0; i < m; i++ {
		Fscan(in, &v)
		if v <= mx {
			Fprint(out, y[sort.Search(len(z), func(i int) bool {
				return z[i] >= v
			})], " ")
		} else if s <= 0 {
			Fprint(out, "-1 ")
		} else {
			cnt := (v - mx + s - 1) / s
			p := sort.Search(len(z), func(i int) bool {
				return z[i] >= v-cnt*s
			})
			Fprint(out, cnt*n+y[p], " ")
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
