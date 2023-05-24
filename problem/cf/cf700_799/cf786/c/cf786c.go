package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func solve(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		Fscan(in, &a[i])
	}

	ans := make([]int, n+1)
	clock := 0
	time := make([]int, n+1)
	f := func(k int) int {
		clock++
		m := 1
		cnt := 0
		for _, v := range a {
			if time[v] == clock {
				continue
			}
			cnt += 1
			if cnt > k {
				m++
				clock++
				cnt = 1
			}
			time[v] = clock
		}
		return m
	}
	var dfs func(l, r int)
	dfs = func(l, r int) {
		if ans[l] == 0 {
			ans[l] = f(l)
		}
		if ans[r] == 0 {
			ans[r] = f(r)
		}
		if l+1 >= r {
			return
		}
		if ans[l] == ans[r] {
			for i := l + 1; i < r; i++ {
				ans[i] = ans[l]
			}
			return
		}
		mid := (l + r) >> 1
		dfs(l, mid)
		dfs(mid, r)
	}
	dfs(1, n)
	for _, v := range ans[1:] {
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
