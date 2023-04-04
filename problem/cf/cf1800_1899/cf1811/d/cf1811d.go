package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

type key struct{ n, x, y int }

func solve(in io.Reader, out io.Writer, fib []int) {
	var n, x, y int
	Fscan(in, &n, &x, &y)
	mp := map[key]bool{}
	var f func(n, x, y int) bool
	f = func(n, x, y int) bool {
		if n == 1 {
			return true
		}

		if v, ok := mp[key{n, x, y}]; ok {
			return v
		}
		h, w := fib[n], fib[n+1]
		ans := false
		if y > h {
			ans = f(n-1, y-h, x)
		}
		if y <= w-h {
			ans = f(n-1, y, x)
		}
		mp[key{n, x, y}] = ans
		return ans
	}
	if f(n, x, y) {
		Fprintln(out, "YES")
	} else {
		Fprintln(out, "NO")
	}
}
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)

	defer out.Flush()
	fib := []int{1, 1}
	for i := 0; i < 44; i++ {
		fib = append(fib, fib[len(fib)-1]+fib[len(fib)-2])
	}
	var t int
	for Fscan(in, &t); t > 0; t-- {
		solve(in, out, fib)
	}
}

func main() { run(os.Stdin, os.Stdout) }
