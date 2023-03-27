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
	f := func(a []int) bool {
		var st, pre []int
		s := 0
		for _, v := range a {
			for len(st) > 0 && st[len(st)-1] <= v {
				if pre[len(pre)-1] < s {
					return false
				}
				st = st[:len(st)-1]
				pre = pre[:len(pre)-1]
			}
			st = append(st, v)
			pre = append(pre, s)
			s += v
		}
		return true
	}
	if !f(a) {
		Fprintln(out, "NO")
		return
	}
	l, r := 0, n-1
	for l < r {
		a[l], a[r] = a[r], a[l]
		l++
		r--
	}
	if !f(a) {
		Fprintln(out, "NO")
		return
	}
	Fprintln(out, "YES")

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
