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
	var n, ans int
	var s string
	cnt := 1
	Fscan(in, &s)
	s = ")" + s
	n = len(s)
	f := make([]int, n)
	for i := 2; i < n; i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				f[i] = f[i-2] + 2
			} else if s[i-1] > 0 && s[i-f[i-1]-1] == '(' {
				f[i] = f[i-1] + 2 + f[i-f[i-1]-2]
			}
			if ans < f[i] {
				ans, cnt = f[i], 1
			} else if ans == f[i] && ans > 0 {
				cnt++
			}
		}
	}

	Fprintln(out, ans, cnt)
}

func main() { run(os.Stdin, os.Stdout) }
