package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func solve(in io.Reader, out io.Writer) {
	var n, k, ans, j int
	Fscan(in, &n, &k)
	a := make([]int, n)
	w := make([]int, k)
	for i := 0; i < n; i++ {
		Fscan(in, &a[i])
	}
	for i := 0; i < k; i++ {
		Fscan(in, &w[i])
	}
	sort.Ints(a)
	sort.Ints(w)
	for i := 0; i < k; i++ {
		if w[i] == 1 {
			ans += a[n-i-1] * 2
		} else {
			ans += a[n-i-1]
		}
	}
	for i := k - 1; i >= 0 && w[i] > 1; i-- {
		ans += a[j]
		j += w[i] - 1
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
