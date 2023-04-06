package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func solve(in io.Reader, out io.Writer) {
	var n int
	var a []string
	var s string
	Fscan(in, &n)
	for i := 0; i < n; i++ {
		Fscan(in, &s)
		a = append(a, s)
	}
	p := map[string]int{}
	p32 := map[string]int{}
	for _, s := range a {
		if s[0] == s[len(s)-1] {
			Fprintln(out, "YES")
			return
		}
		if len(s) == 2 {
			sb := make([]byte, 2)
			sb[0], sb[1] = s[1], s[0]
			t := string(sb)
			if p[t] == 1 || p32[t] == 1 {
				Fprintln(out, "YES")
				return
			}
			p[s] = 1
		} else if len(s) == 3 {
			sb := make([]byte, 3)
			sb[0], sb[1], sb[2] = s[2], s[1], s[0]
			if p[string(sb)] == 1 || p[string(sb[:2])] == 1 {
				Fprintln(out, "YES")
				return
			}
			p[s] = 1
			p32[s[:2]] = 1
		}
	}

	Fprintln(out, "NO")
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
