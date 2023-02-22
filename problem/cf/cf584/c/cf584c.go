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
	var n, t int
	var s1, s2 string
	Fscan(in, &n, &t, &s1, &s2)
	p := n - t
	var same []int
	var diff []int
	for i := 0; i < n; i++ {
		if s1[i] == s2[i] {
			same = append(same, i)
		} else {
			diff = append(diff, i)
		}
	}
	s := len(same)
	if p > s && (p-s)*2 > len(diff) {
		Fprintln(out, -1)
		return
	}
	ans := make([]byte, n)
	if p <= s {
		for _, i := range same[:p] {
			ans[i] = s1[i]
		}
	} else {
		for _, i := range same {
			ans[i] = s1[i]
		}
		for i := 0; i < p-s; i++ {
			ans[diff[i*2]] = s1[diff[i*2]]
			ans[diff[i*2+1]] = s2[diff[i*2+1]]
		}
	}

	for i, c := range ans {
		if c == 0 {
			for d := byte('a'); d <= byte('z'); d++ {
				if s1[i] != d && s2[i] != d {
					ans[i] = d
					break
				}
			}
		}
	}
	Fprintln(out, string(ans))

}

func main() { run(os.Stdin, os.Stdout) }
