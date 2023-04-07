package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func solve(in io.Reader, out io.Writer) {
	var n, ans int
	Fscan(in, &n)
	a := make([]int, n)
	cnt := map[int]int{}
	for i := 0; i < n; i++ {
		Fscan(in, &a[i])
		cnt[a[i]]++
	}
	ans = 2
	for mn, keep := range cnt {
		mx := mn
		p := sort.Search(n, func(i int) bool {
			return a[i] >= max(mx*2-mn, mn+1)
		})
		for p < n {
			mx = a[p]
			keep++
			p = sort.Search(n, func(i int) bool {
				return a[i] >= max(mx*2-mn, mn+1)
			})
		}
		ans = max(ans, keep)
	}

	Fprintln(out, n-ans)
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
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
