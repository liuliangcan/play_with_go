package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func solve(in io.Reader, out io.Writer) {
	var n, mx, k int
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		Fscan(in, &a[i])
		mx = max(mx, a[i])
	}
	ans := sort.Search(mx+1, func(x int) bool {
		f := func(i int) bool {
			cnt := i
			for i < n {
				if a[i] <= x {
					cnt++
					i++
					if i < n {
						cnt++
						i++
					}
					continue
				}
				i++
				if cnt >= k {
					return true
				}
			}
			return cnt >= k
		}
		return f(0) || f(1)
	})

	Fprintln(out, ans)
}
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	solve(in, out)
}

func main() { run(os.Stdin, os.Stdout) }
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
