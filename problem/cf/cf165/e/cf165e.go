package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		Fscan(in, &a[i])
	}
	const INF = (1 << 22) - 1
	f := make([]int, INF+1)
	for i := range f {
		f[i] = -1
	}
	for _, v := range a {
		f[(v^INF)&INF] = v
	}
	for i := INF - 1; i > 0; i-- {
		if f[i] == -1 {
			for j := 0; j < 22; j++ {
				if (i>>j)&1 == 0 {
					p := i | (1 << j)
					if f[p] != -1 {
						f[i] = f[p]
						break
					}
				}
			}
		}
	}
	for _, v := range a {
		Fprint(out, f[v], " ")
	}

}

func main() { run(os.Stdin, os.Stdout) }
