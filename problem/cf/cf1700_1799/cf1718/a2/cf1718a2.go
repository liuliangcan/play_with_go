package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func solve(in io.Reader, out io.Writer) {
	var n, v, xor int
	Fscan(in, &n)
	ans := n
	has := map[int]bool{0: true}
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		xor ^= v
		if has[xor] {
			ans -= 1
			has = map[int]bool{xor: true}
		} else {
			has[xor] = true
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
