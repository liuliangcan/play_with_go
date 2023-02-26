package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func solve(in io.Reader, out io.Writer) {
	var n, ans int
	Fscan(in, &n)
	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		Fscan(in, &a[i])
	}
	for i := 0; i < n; i++ {
		Fscan(in, &b[i])
	}
	ok := func(mask int) bool {
		cnt := map[int]int{}
		for _, v := range a {
			cnt[v&mask]++
		}
		for _, v := range b {
			cnt[(^v)&mask]--
		}
		for _, v := range cnt {
			if v != 0 {
				return false
			}
		}
		return true
	}
	for i := 29; i >= 0; i-- {
		if ok(ans | (1 << i)) {
			ans |= 1 << i
		}
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
