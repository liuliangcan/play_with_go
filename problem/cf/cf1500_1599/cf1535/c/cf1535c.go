package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func solve(in io.Reader, out io.Writer) {
	var ans int
	var s string
	Fscan(in, &s)
	pos := []int{-1, -1}
	for i, c := range s {
		if c != '?' {
			pos[(int(c)^i)&1] = i
		}
		ans += i - min(pos[0], pos[1])
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
	if b < a {
		return b
	}
	return a
}
