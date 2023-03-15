package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func solve(in io.Reader, out io.Writer) {
	var n, k, v int
	Fscan(in, &n, &k)
	f := make([][]int, k+1)
	for i := range f {
		f[i] = make([]int, k+1)
	}
	f[0][0] = 1
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		for j1 := k; j1 >= 0; j1-- {
			for j2 := k; j2 >= 0; j2-- {
				if j1 >= v {
					f[j1][j2] |= f[j1-v][j2]
				}
				if j2 >= v {
					f[j1][j2] |= f[j1][j2-v]
				}
			}
		}
	}
	var ans []int
	for i, g := range f {
		if g[k-i] == 1 {
			ans = append(ans, i)
		}
	}

	Fprintln(out, len(ans))
	for _, v := range ans {
		Fprint(out, v, " ")
	}
}
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	solve(in, out)
}

func main() { run(os.Stdin, os.Stdout) }
