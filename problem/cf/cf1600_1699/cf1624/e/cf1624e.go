package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func solve(in io.Reader, out io.Writer) {
	var n, m int
	var s string
	Fscan(in, &n, &m)
	type pos struct{ l, r, i int }
	two := map[string]pos{}
	three := map[string]pos{}
	for i := 1; i <= n; i++ {
		Fscan(in, &s)
		if m == 1 {
			continue
		}
		two[s[0:2]] = pos{1, 2, i}
		for j := 3; j <= m; j++ {
			two[s[j-2:j]] = pos{j - 1, j, i}
			three[s[j-3:j]] = pos{j - 2, j, i}
		}
	}
	Fscan(in, &s)
	if m == 1 {
		Fprintln(out, "-1")
		return
	}
	f := make([]bool, m+1)
	f[0] = true
	if two[s[:2]].l > 0 {
		f[2] = true
	}
	for j := 3; j <= m; j++ {
		f[j] = f[j-2] && two[s[j-2:j]].l > 0 || f[j-3] && three[s[j-3:j]].l > 0
	}
	if !f[m] {
		Fprintln(out, "-1")
		return
	}
	var ans []pos
	j := m
	for j > 0 {
		if f[j-2] && two[s[j-2:j]].l > 0 {
			ans = append(ans, two[s[j-2:j]])
			j -= 2
		} else {
			ans = append(ans, three[s[j-3:j]])
			j -= 3
		}
	}

	Fprintln(out, len(ans))
	for j := len(ans) - 1; j >= 0; j-- {
		Fprintln(out, ans[j].l, ans[j].r, ans[j].i)
	}
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
