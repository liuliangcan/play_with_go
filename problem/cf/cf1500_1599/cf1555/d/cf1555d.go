package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func solve(in io.Reader, out io.Writer) {
	var n, m, l, r int
	var s []byte
	Fscan(in, &n, &m, &s)
	perm3 := []string{
		"abc",
		"acb",
		"bac",
		"bca",
		"cab",
		"cba",
	}

	pre := make([][]int, 6)
	for i, t := range perm3 {
		pre[i] = make([]int, n+1)
		for j, c := range s {
			pre[i][j+1] = pre[i][j]
			if t[j%3] != c {
				pre[i][j+1]++
			}
		}
	}

	for ; m > 0; m-- {
		Fscan(in, &l, &r)
		ans := n
		for _, p := range pre {
			ans = min(ans, p[r]-p[l-1])
		}
		Fprintln(out, ans)
	}

}
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	solve(in, out)
}

func main() { run(os.Stdin, os.Stdout) }
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
