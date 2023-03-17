package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func solve(in io.Reader, out io.Writer) {
	var n, mx int
	Fscan(in, &n)
	h := make([]int, n)

	for i := 0; i < n; i++ {
		Fscan(in, &h[i])
		if h[i] > mx {
			mx = h[i]
		}
	}
	f := func(t int) int {
		a, b := 0, 0
		for _, v := range h {
			a += (t - v) % 2 // 填1的次数
			b += (t - v) / 2 // 填2的次数
		}
		if a == b {
			return a * 2
		} else if a > b {
			return a*2 - 1
		} else {
			ans := a * 2
			b -= a
			ans += b * 2 / 3 * 2 // 花2天填3个
			ans += b * 2 % 3
			return ans
		}
	}

	Fprintln(out, min(f(mx+1), f(mx)))
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
