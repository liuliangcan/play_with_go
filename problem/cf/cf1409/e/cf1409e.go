package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func solve(in io.Reader, out io.Writer) {
	var n, ans, k int
	Fscan(in, &n, &k)
	a := make([]int, n)
	f := make([]int, n+1)
	for i := 0; i < n; i++ {
		Fscan(in, &a[i])
	}
	for i := 0; i < n; i++ {
		Fscan(in, &f[1])
	}
	sort.Ints(a)
	l := 0
	for i, v := range a {
		for v-a[l] > k {
			l++
		}
		f[i+1] = max(f[i], i-l+1)
		ans = max(ans, i-l+1+f[l])
	}
	Fprintln(out, ans)
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
