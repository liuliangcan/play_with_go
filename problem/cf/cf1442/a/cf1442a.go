package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var t, n int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n)
		a := make([]int, n)
		preMin := make([]int, n)
		for i := 0; i < n; i++ {
			Fscan(in, &a[i])
			if i == 0 {
				preMin[i] = a[i]
			} else {
				preMin[i] = min(preMin[i-1], a[i])
			}
		}
		delta := 0
		for i := n - 2; i >= 0; i-- {
			a[i] -= delta
			if a[i] > a[i+1] {
				d := a[i] - a[i+1]
				if preMin[i]-delta >= d {
					delta += d
					a[i] -= d
				} else {
					delta = -1
					break
				}
			}
		}
		if delta == -1 {
			Fprintln(out, "NO")
		} else {
			Fprintln(out, "YES")
		}
	}

}

func main() { run(os.Stdin, os.Stdout) }
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
