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
	Fscan(in, &n, &s)
	if n&1 == 1 {
		Fprintln(out, 0)
		return
	}
	suf := make([]int, n+1)
	j := n - 1
	for ; j >= 0; j-- {
		if s[j] == ')' {
			suf[j] = suf[j+1] + 1
		} else {
			suf[j] = suf[j+1] - 1
			if suf[j] < 0 {
				break
			}
		}
	}
	p := 0
	for i, c := range s {
		if c == '(' {
			if i >= j && p > 0 && p == suf[i+1]+1 {
				ans += 1
			}
			p++
		} else {
			if i >= j && suf[i+1] > 0 && p+1 == suf[i+1] {
				ans += 1
			}
			p--
			if p < 0 {
				break
			}
		}
	}

	Fprintln(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
