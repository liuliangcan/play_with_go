package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func solve(in io.Reader, out io.Writer) {
	var n, k, s int
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		Fscan(in, &a[i])
		s += a[i]
	}
	if s <= k {
		Fprintln(out, "0")
		return
	}
	if n == 1 {
		Fprintln(out, s-k)
		return
	}
	sort.Ints(a)
	l := a[0]
	d := s - k
	ans := s - k
	cnt := 0
	s = 0
	for i := n - 1; i > 0; i-- {
		s += a[i]
		cnt += 1
		if s-l*cnt >= d {
			ans = min(ans, cnt)
			break
		}
		x := (l*cnt - s + d + cnt + 1 - 1) / (cnt + 1)
		ans = min(ans, x+cnt)
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
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
