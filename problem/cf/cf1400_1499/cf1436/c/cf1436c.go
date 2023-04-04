package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

const MOD = int(1e9 + 7)

func perm(a, b int) int {
	if a < b {
		return 0
	}
	ans := 1
	for i := a; i > a-b; i-- {
		ans = ans * i % MOD
	}
	return ans
}
func solve(in io.Reader, out io.Writer) {
	var n, x, pos int
	Fscan(in, &n, &x, &pos)
	l, r := 0, n
	less, greater := 0, 0
	for l < r {
		mid := (l + r) / 2
		if mid <= pos {
			l = mid + 1
			if mid < pos {
				less++
			}
		} else if mid > pos {
			r = mid
			greater++
		}
	}

	Fprintln(out, perm(x-1, less)*perm(n-x, greater)%MOD*perm(n-less-greater-1, n-less-greater-1)%MOD)
}
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	solve(in, out)
}

func main() { run(os.Stdin, os.Stdout) }
