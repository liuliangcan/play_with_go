package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func solve(in io.Reader, out io.Writer) {
	var n, k int
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		Fscan(in, &a[i])
	}
	sort.Ints(a)

	Fprintln(out, sort.Search(a[n-1]-a[0], func(x int) bool {
		f := make([]bool, n+1)
		f[0] = true
		for i, j := k-1, 0; i < n; i++ {
			for a[i]-a[j] > x {
				j++
			}
			for ; j <= i-k+1; j++ {
				if f[j] {
					f[i+1] = true
					break
				}
			}
		}
		return f[n]
	}))
}
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	solve(in, out)
}

func main() { run(os.Stdin, os.Stdout) }
